# Cadillac
(*Cadillac*)

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
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2"
	"github.com/smartcar/go-sdk-v2/pkg/models/operations"
	"context"
	"log"
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
    request := operations.GetCadillacChargeTimeRequest{
        VehicleID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Cadillac.GetChargeTime(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ChargeTime != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCadillacChargeTimeRequest](../../pkg/models/operations/getcadillacchargetimerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.GetCadillacChargeTimeResponse](../../pkg/models/operations/getcadillacchargetimeresponse.md), error**
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
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2"
	"github.com/smartcar/go-sdk-v2/pkg/models/operations"
	"context"
	"log"
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
    request := operations.GetCadillacVoltageRequest{
        VehicleID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Cadillac.GetVoltage(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ChargeVoltage != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetCadillacVoltageRequest](../../pkg/models/operations/getcadillacvoltagerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.GetCadillacVoltageResponse](../../pkg/models/operations/getcadillacvoltageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
