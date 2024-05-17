/*
Cart Service

Testing CartAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cart

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Gemini-Commerce/go-client-cart"
)

func Test_cart_CartAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CartAPIService CartSetCustomPriceOnItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CartAPI.CartSetCustomPriceOnItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}