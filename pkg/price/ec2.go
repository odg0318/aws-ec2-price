package price

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	URL = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonEC2/current/index.json"

	HOURLY_TERM_CODE = "JRTCKXETXF"
	RATE_CODE        = "6YS6EN2CT7"

	REQUIRED_TERM           = "OnDemand"
	REQUIRED_TENANCY        = "Shared"
	REQUIRED_PRODUCT_FAMILY = "Compute Instance"
	REQUIRED_OS             = "Linux"
	REQUIRED_LICENSE_MODEL  = "No License required"

	CACHED_PRICING = CachedEc2Pricing{}
)

type Ec2Pricing struct {
	Products map[string]Ec2Product `json:products`
	Terms    map[string]map[string]map[string]struct {
		PriceDimensions map[string]struct {
			PricePerUnit struct {
				USD string `json:USD`
			} `json:pricePerUnit`
		} `json:priceDimensions`
	} `json:terms`
}

func (ec *Ec2Pricing) GetPrice(region string, instanceType string) (string, error) {
	for _, product := range ec.Products {
		if product.isValid() == false {
			continue
		}

		if product.isValidRegion(region) == false {
			continue
		}

		if product.isValidInstanceType(instanceType) == false {
			continue
		}

		h := fmt.Sprintf("%s.%s", product.Sku, HOURLY_TERM_CODE)
		r := fmt.Sprintf("%s.%s.%s", product.Sku, HOURLY_TERM_CODE, RATE_CODE)

		price := ec.Terms[REQUIRED_TERM][product.Sku][h].PriceDimensions[r].PricePerUnit.USD

		return price, nil
	}

	return "", errors.New("no price")
}

type Ec2Product struct {
	Sku           string `json:sku`
	ProductFamily string `json:productFamily`
	Attributes    struct {
		Location        string `json:location`
		InstanceType    string `json:instanceType`
		Tenancy         string `json:tenancy`
		OperatingSystem string `json:operatingSystem`
		LicenseModel    string `json:licenseModel`
	}
}

func (ep *Ec2Product) isValidRegion(region string) bool {
	if r, ok := REGIONS[region]; ok == true {
		return ep.Attributes.Location == r
	}

	return false
}

func (ep *Ec2Product) isValidInstanceType(instanceType string) bool {
	return ep.Attributes.InstanceType == instanceType
}

func (ep *Ec2Product) isValid() bool {
	if ep.ProductFamily != REQUIRED_PRODUCT_FAMILY {
		return false
	}

	if ep.Attributes.OperatingSystem != REQUIRED_OS {
		return false
	}

	if ep.Attributes.LicenseModel != REQUIRED_LICENSE_MODEL {
		return false
	}

	if ep.Attributes.Tenancy != REQUIRED_TENANCY {
		return false
	}

	return true

}

type CachedEc2Pricing struct {
	pricing       *Ec2Pricing
	lastCheckTime time.Time
}

func (c *CachedEc2Pricing) isExpired() bool {
	return time.Since(c.lastCheckTime) > time.Duration(24*time.Hour)
}

func (c *CachedEc2Pricing) update(pricing *Ec2Pricing) {
	c.pricing = pricing
	c.lastCheckTime = time.Now()
}

func NewPricing() (*Ec2Pricing, error) {
	if CACHED_PRICING.isExpired() == false {
		return CACHED_PRICING.pricing, nil
	}

	client := &http.Client{}

	r, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	pricing := &Ec2Pricing{}
	if err := json.NewDecoder(r.Body).Decode(pricing); err != nil {
		return nil, err
	}

	CACHED_PRICING.update(pricing)

	return pricing, err
}
