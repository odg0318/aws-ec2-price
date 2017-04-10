package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/odg0318/aws-ec2-price/pkg/price"
	"github.com/stretchr/testify/assert"
)

func TestGetEc2PricesHandler(t *testing.T) {
	r := GetRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	url := fmt.Sprintf("%s/ec2/regions/%s", server.URL, "us-east-1")
	response, err := http.Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)

	instances := []*price.Instance{}
	err = json.Unmarshal(body, &instances)

	assert.True(t, len(instances) > 0)

	for _, i := range instances {
		assert.Equal(t, "us-east-1", i.Region)
	}
}
