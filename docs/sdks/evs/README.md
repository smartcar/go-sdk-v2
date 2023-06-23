# Evs

## Overview

Operations about electric vehicles

### Available Operations

* [GetBatteryCapacity](#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](#getbatterylevel) - EV Battery Level
* [GetChargingLimit](#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](#getchargingstatus) - EV Charging Status
* [SetChargingLimit](#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](#startstopcharge) - Start or stop charging an electric vehicle. Please contact us to request early access.

## GetBatteryCapacity

__Description__

Returns the total capacity of an electric vehicle's battery.

__Permission__

`read_battery`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  capacity|   number|  The total capacity of the vehicle's battery (in kilowatt-hours). 	|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"smartcar"
	"smartcar/pkg/models/operations"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetBatteryCapacity(ctx, operations.GetBatteryCapacityRequest{
        VehicleID: "unde",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatteryCapacity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetBatteryCapacityRequest](../../models/operations/getbatterycapacityrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetBatteryCapacityResponse](../../models/operations/getbatterycapacityresponse.md), error**


## GetBatteryLevel

__Description__

Retrieve EV battery level of a vehicle.

__Permission__

`read_battery`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `percentRemaining`|   number|  An EV battery’s state of charge (in percent). 	|
|   `range`|   number	|   The estimated remaining distance the vehicle can travel (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers).	|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"smartcar"
	"smartcar/pkg/models/operations"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetBatteryLevel(ctx, operations.GetBatteryLevelRequest{
        VehicleID: "nulla",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatteryLevel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetBatteryLevelRequest](../../models/operations/getbatterylevelrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetBatteryLevelResponse](../../models/operations/getbatterylevelresponse.md), error**


## GetChargingLimit

__Description__

Returns the current charge limit of an electric vehicle.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"smartcar"
	"smartcar/pkg/models/operations"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetChargingLimit(ctx, operations.GetChargingLimitRequest{
        VehicleID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeLimit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetChargingLimitRequest](../../models/operations/getcharginglimitrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetChargingLimitResponse](../../models/operations/getcharginglimitresponse.md), error**


## GetChargingStatus

__Description__

Returns the current charge status of an electric vehicle.

__Permission__

`read_charge`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `isPluggedIn` 	|   boolean	|  Indicates whether a charging cable is currently plugged into the vehicle’s charge port. 	|
|   `state`	|   string	|   Indicates whether the vehicle is currently charging. Options: `CHARGING` `FULLY_CHARGED` `NOT_CHARGING`	|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"smartcar"
	"smartcar/pkg/models/operations"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetChargingStatus(ctx, operations.GetChargingStatusRequest{
        VehicleID: "illum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetChargingStatusRequest](../../models/operations/getchargingstatusrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetChargingStatusResponse](../../models/operations/getchargingstatusresponse.md), error**


## SetChargingLimit

__Description__

Returns the current charge limit of an electric vehicle.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"smartcar"
	"smartcar/pkg/models/operations"
	"smartcar/pkg/models/shared"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.SetChargingLimit(ctx, operations.SetChargingLimitRequest{
        ChargeLimit: &shared.ChargeLimit{
            Limit: smartcar.Float32(1),
        },
        VehicleID: "vel",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.SetChargingLimitRequest](../../models/operations/setcharginglimitrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.SetChargingLimitResponse](../../models/operations/setcharginglimitresponse.md), error**


## StartStopCharge

__Description__

Returns the current charge status of an electric vehicle.

__Permission__

`read_charge`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `isPluggedIn` 	|   boolean	|  Indicates whether a charging cable is currently plugged into the vehicle’s charge port. 	|
|   `state`	|   string	|   Indicates whether the vehicle is currently charging. Options: `CHARGING` `FULLY_CHARGED` `NOT_CHARGING`	|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"smartcar"
	"smartcar/pkg/models/operations"
	"smartcar/pkg/models/shared"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.StartStopCharge(ctx, operations.StartStopChargeRequest{
        ChargeAction: &shared.ChargeAction{
            Action: shared.ChargeActionActionStart.ToPointer(),
        },
        VehicleID: "error",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.StartStopChargeRequest](../../models/operations/startstopchargerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.StartStopChargeResponse](../../models/operations/startstopchargeresponse.md), error**

