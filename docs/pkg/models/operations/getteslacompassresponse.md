# GetTeslaCompassResponse


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Compass`                                                | [*shared.Compass](../../../pkg/models/shared/compass.md) | :heavy_minus_sign:                                       | returns the compass heading of a Tesla.                  |
| `ContentType`                                            | *string*                                                 | :heavy_check_mark:                                       | HTTP response content type for this operation            |
| `StatusCode`                                             | *int*                                                    | :heavy_check_mark:                                       | HTTP response status code for this operation             |
| `RawResponse`                                            | [*http.Response](https://pkg.go.dev/net/http#Response)   | :heavy_check_mark:                                       | Raw HTTP response; suitable for custom response parsing  |