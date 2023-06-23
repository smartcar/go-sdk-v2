# GetChargingLimitResponse


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ChargeLimit`                                             | [*shared.ChargeLimit](../../models/shared/chargelimit.md) | :heavy_minus_sign:                                        | return EV Charge Limit                                    |
| `ContentType`                                             | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `StatusCode`                                              | *int*                                                     | :heavy_check_mark:                                        | N/A                                                       |
| `RawResponse`                                             | [*http.Response](https://pkg.go.dev/net/http#Response)    | :heavy_minus_sign:                                        | N/A                                                       |