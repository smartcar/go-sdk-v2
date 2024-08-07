# GetBatteryLevelResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `BatteryLevel`                                                     | [*shared.BatteryLevel](../../../pkg/models/shared/batterylevel.md) | :heavy_minus_sign:                                                 | return EV Battery Level reading                                    |
| `ContentType`                                                      | *string*                                                           | :heavy_check_mark:                                                 | HTTP response content type for this operation                      |
| `StatusCode`                                                       | *int*                                                              | :heavy_check_mark:                                                 | HTTP response status code for this operation                       |
| `RawResponse`                                                      | [*http.Response](https://pkg.go.dev/net/http#Response)             | :heavy_check_mark:                                                 | Raw HTTP response; suitable for custom response parsing            |