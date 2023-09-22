# Cadillac

### Available Operations

* [GetChargeTime](#getchargetime) - Retrieve charging completion time for a Cadillac.
* [GetVoltage](#getvoltage) - Retrieve charging voltmeter time for a Cadillac.

## GetChargeTime

__Description__

When the vehicle is charging, this endpoint returns the date and time the vehicle expects to complete this charging session. When the vehicle is not charging, this endpoint results in a vehicle state error.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkv2 "github.com/smartcar/go-sdk-v2"
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"github.com/smartcar/go-sdk-v2/pkg/models/operations"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Cadillac.GetChargeTime(ctx, operations.GetCadillacChargeTimeRequest{
        VehicleID: "corrupti",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetCadillacChargeTimeRequest](../../models/operations/getcadillacchargetimerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetCadillacChargeTimeResponse](../../models/operations/getcadillacchargetimeresponse.md), error**


## GetVoltage

__Description__

When the vehicle is plugged in, this endpoint returns the voltage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkv2 "github.com/smartcar/go-sdk-v2"
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"github.com/smartcar/go-sdk-v2/pkg/models/operations"
)

func main() {
    s := gosdkv2.New(
        gosdkv2.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Cadillac.GetVoltage(ctx, operations.GetCadillacVoltageRequest{
        VehicleID: "provident",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCadillacVoltageRequest](../../models/operations/getcadillacvoltagerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetCadillacVoltageResponse](../../models/operations/getcadillacvoltageresponse.md), error**

