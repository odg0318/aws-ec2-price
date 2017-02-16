package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	EXPECTED_REGION        = "us-east-1"
	EXPECTED_INSTANCE_TYPE = "c4.large"
)

func TestGetEc2PriceHandler(t *testing.T) {
	r := GetRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	url := fmt.Sprintf("%s/ec2/regions/%s/instance_types/%s", server.URL, EXPECTED_REGION, EXPECTED_INSTANCE_TYPE)
	response, err := http.Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)

	price := map[string]string{}
	err = json.Unmarshal(body, &price)

	assert.Equal(t, EXPECTED_REGION, price["region"])
	assert.Equal(t, EXPECTED_INSTANCE_TYPE, price["instance_type"])
}

func TestGetEc2PriceHandlerWithInvalid(t *testing.T) {
	r := GetRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	url := fmt.Sprintf("%s/ec2/regions/%s/instance_types/%s", server.URL, "invalid-region", "invalid-instance-type")
	response, err := http.Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)
}
