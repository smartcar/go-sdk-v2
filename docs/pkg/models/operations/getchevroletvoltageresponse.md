# GetChevroletVoltageResponse


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ChargeVoltage`                                                      | [*shared.ChargeVoltage](../../../pkg/models/shared/chargevoltage.md) | :heavy_minus_sign:                                                   | returns the voltage of the charger measured by the vehicle.          |
| `ContentType`                                                        | *string*                                                             | :heavy_check_mark:                                                   | HTTP response content type for this operation                        |
| `StatusCode`                                                         | *int*                                                                | :heavy_check_mark:                                                   | HTTP response status code for this operation                         |
| `RawResponse`                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)               | :heavy_check_mark:                                                   | Raw HTTP response; suitable for custom response parsing              |