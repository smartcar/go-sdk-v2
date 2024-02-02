# Tesla
(*Tesla*)

### Available Operations

* [GetAmmeter](#getammeter) - Retrieve charging ammeter time for a Tesla.
* [GetChargeTime](#getchargetime) - Retrieve charging completion time for a Tesla.
* [GetCompass](#getcompass) - Retrieve compass heading for a Tesla.
* [GetExteriorTemperature](#getexteriortemperature) - Retrieve exterior temperature for a Tesla.
* [GetInteriorTemperature](#getinteriortemperature) - Retrieve interior temperature for a Tesla.
* [GetSpeedometer](#getspeedometer) - Retrieve speed for a Tesla.
* [GetVoltage](#getvoltage) - Retrieve charging voltmeter time for a Tesla.
* [GetWattmeter](#getwattmeter) - Retrieve charging wattmeter time for a Tesla.
* [SetAmmeter](#setammeter) - Set charging ammeter time for a Tesla.

## GetAmmeter

__Description__

When the vehicle is plugged in, this endpoint returns the amperage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetAmmeter(ctx, operations.GetTeslaAmmeterRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeAmmeter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTeslaAmmeterRequest](../../pkg/models/operations/getteslaammeterrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetTeslaAmmeterResponse](../../pkg/models/operations/getteslaammeterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetChargeTime

__Description__

When the vehicle is charging, this endpoint returns the date and time the vehicle expects to complete this charging session. When the vehicle is not charging, this endpoint results in a vehicle state error.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetChargeTime(ctx, operations.GetTeslaChargeTimeRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeTime != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetTeslaChargeTimeRequest](../../pkg/models/operations/getteslachargetimerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetTeslaChargeTimeResponse](../../pkg/models/operations/getteslachargetimeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCompass

__Description__

This endpoint returns the compass heading of a Tesla. The value is in degrees, with 0 degrees being North.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetCompass(ctx, operations.GetTeslaCompassRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Compass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTeslaCompassRequest](../../pkg/models/operations/getteslacompassrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetTeslaCompassResponse](../../pkg/models/operations/getteslacompassresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetExteriorTemperature

__Description__

This endpoint returns the exterior temperature of a Tesla, in celsius by default.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetExteriorTemperature(ctx, operations.GetTeslaExteriorTemperatureRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Temperature != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetTeslaExteriorTemperatureRequest](../../pkg/models/operations/getteslaexteriortemperaturerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetTeslaExteriorTemperatureResponse](../../pkg/models/operations/getteslaexteriortemperatureresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetInteriorTemperature

__Description__

This endpoint returns the interior temperature of a Tesla, in celsius by default.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetInteriorTemperature(ctx, operations.GetTeslaInteriorTemperatureRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Temperature != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetTeslaInteriorTemperatureRequest](../../pkg/models/operations/getteslainteriortemperaturerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetTeslaInteriorTemperatureResponse](../../pkg/models/operations/getteslainteriortemperatureresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSpeedometer

__Description__

This endpoint returns the speed of a Tesla (in kilometers/hour by default or in miles/hour using the sc-unit-system).

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetSpeedometer(ctx, operations.GetTeslaSpeedRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Speed != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetTeslaSpeedRequest](../../pkg/models/operations/getteslaspeedrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetTeslaSpeedResponse](../../pkg/models/operations/getteslaspeedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetVoltage

__Description__

When the vehicle is plugged in, this endpoint returns the voltage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetVoltage(ctx, operations.GetTeslaVoltageRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeVoltage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTeslaVoltageRequest](../../pkg/models/operations/getteslavoltagerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetTeslaVoltageResponse](../../pkg/models/operations/getteslavoltageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetWattmeter

__Description__

When the vehicle is plugged in, this endpoint returns the wattage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.GetWattmeter(ctx, operations.GetTeslaWattmeterRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeWattage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetTeslaWattmeterRequest](../../pkg/models/operations/getteslawattmeterrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetTeslaWattmeterResponse](../../pkg/models/operations/getteslawattmeterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SetAmmeter

__Description__

When the vehicle is plugged in, this endpoint sets the amperage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Tesla.SetAmmeter(ctx, operations.SetTeslaAmmeterRequest{
        ChargeAmmeter: &shared.ChargeAmmeter{
            Amperage: gosdkv2.Float32(48),
        },
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SetTeslaAmmeterRequest](../../pkg/models/operations/setteslaammeterrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SetTeslaAmmeterResponse](../../pkg/models/operations/setteslaammeterresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
