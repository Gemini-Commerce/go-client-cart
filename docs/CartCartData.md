# # CartCartData


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** |   | [optional]
**Grn**| **string** |   | [optional]
**Channel**| **string** |   | [optional]
**Market**| **string** |   | [optional]
**Locale**| **string** |   | [optional]
**Items**| [**[]CartCartItem**](CartCartItem.md) |   | [optional]
**Payments**| [**[]CartPaymentData**](CartPaymentData.md) |   | [optional]
**Shipments**| [**[]CartShipmentData**](CartShipmentData.md) |   | [optional]
**Promotions**| [**[]CartPromotionData**](CartPromotionData.md) |   | [optional]
**Currency**| [**CartCurrency**](CartCurrency.md) |  for more information please, see Model/CartCurrency.php  | [optional] [default to CARTCURRENCY_XXX]
**Subtotals**| [**[]CartCartSubtotal**](CartCartSubtotal.md) |   | [optional]
**Total**| [**CartMoney**](CartMoney.md) |   | [optional]
**TotalDue**| [**CartMoney**](CartMoney.md) |   | [optional]
**VatIncluded**| **bool** |   | [optional]
**BillingAddress**| [**CartPostalAddress**](CartPostalAddress.md) |   | [optional]
**ShippingAddress**| [**CartPostalAddress**](CartPostalAddress.md) |   | [optional]
**Status**| [**CartCartStatus**](CartCartStatus.md) |  for more information please, see Model/CartCartStatus.php  | [optional] [default to CARTCARTSTATUS_UNKNOWN]
**Customer**| [**CartCustomerData**](CartCustomerData.md) |   | [optional]
**Notes**| **string** |   | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

