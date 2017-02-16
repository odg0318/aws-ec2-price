package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoRouteHandler(t *testing.T) {
	r := GetRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	url := fmt.Sprintf("%s/noroute", server.URL)
	response, err := http.Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 404, response.StatusCode)
}
