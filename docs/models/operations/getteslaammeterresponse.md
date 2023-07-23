# GetTeslaAmmeterResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ChargeAmmeter`                                               | [*shared.ChargeAmmeter](../../models/shared/chargeammeter.md) | :heavy_minus_sign:                                            | returns the amperage of the charger measured by the vehicle.  |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |