# GetBatteryCapacityResponse


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BatteryCapacity`                                                 | [*shared.BatteryCapacity](../../models/shared/batterycapacity.md) | :heavy_minus_sign:                                                | return EV Battery Capacity reading                                |
| `ContentType`                                                     | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `StatusCode`                                                      | *int*                                                             | :heavy_check_mark:                                                | N/A                                                               |
| `RawResponse`                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)            | :heavy_minus_sign:                                                | N/A                                                               |