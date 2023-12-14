# Vehicles
(*Vehicles*)

## Overview

Operations about vehicles

### Available Operations

* [Batch](#batch) - Batch
* [LockStatus](#lockstatus) - Returns the lock status for a vehicle and the open status of its doors, windows, storage units, sunroof and charging port where available. The open status array(s) will be empty if a vehicle has partial support. The request will error if lock status can not be retrieved from the vehicle or the brand is not supported.
* [Disconnect](#disconnect) - Revoke Access
* [Get](#get) - Vehicle Info
* [GetEngineOil](#getengineoil) - Engine Oil Life
* [GetFuelTank](#getfueltank) - Fuel Tank (US Only)
* [GetLocation](#getlocation) - Location
* [GetOdometer](#getodometer) - Odometer
* [GetPermissions](#getpermissions) - Application Permissions
* [GetTirePressure](#gettirepressure) - Tire Pressure
* [GetVin](#getvin) - Returns the vehicle’s manufacturer identifier.
* [ListVehicles](#listvehicles) - All Vehicles
* [LockUnlock](#lockunlock) - Lock/Unlock Vehicle
* [SendDestination](#senddestination) - Send Destination to vehicle navigation system

## Batch

__Description__ Returns a list of responses from multiple Smartcar endpoints, all combined into a single request. Note: Batch requests is a paid feature. Please contact us to upgrade your plan and obtain access.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.Batch(ctx, operations.BatchRequest{
        RequestBody: []string{
            "/odometer",
        },
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.BatchRequest](../../pkg/models/operations/batchrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.BatchResponse](../../pkg/models/operations/batchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## LockStatus

Returns the lock status for a vehicle and the open status of its doors, windows, storage units, sunroof and charging port where available. The open status array(s) will be empty if a vehicle has partial support. The request will error if lock status can not be retrieved from the vehicle or the brand is not supported.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.LockStatus(ctx, operations.LockStatusRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SecurityRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.LockStatusRequest](../../pkg/models/operations/lockstatusrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.LockStatusResponse](../../pkg/models/operations/lockstatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Disconnect

__Description__

Revoke access for the current requesting application. This is the correct way to disconnect a vehicle.

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.Disconnect(ctx, operations.DisconnectRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Status != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DisconnectRequest](../../pkg/models/operations/disconnectrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DisconnectResponse](../../pkg/models/operations/disconnectresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Get

__Permission__

Returns a single vehicle object, containing identifying information.

__Permission__

`read_vehicle_info`

__Response Body__

|Name| Type|Description|
|--|--|--|
|`id`|string|A vehicle ID (UUID v4).|
|`make`|string|The manufacturer of the vehicle.|
|`model`|integer|The model of the vehicle.|
|`year`|integer|The model year.|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.Get(ctx, operations.GetVehicleRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VehicleInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetVehicleRequest](../../pkg/models/operations/getvehiclerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetVehicleResponse](../../pkg/models/operations/getvehicleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetEngineOil

__Description__

Returns the remaining life span of a vehicle’s engine oil.

__Permission__

`read_engine_oil`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `lifeRemaining`|   number|  The engine oil’s remaining life span (as a percentage). Oil life is based on the current quality of the oil. (in percent).|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.GetEngineOil(ctx, operations.GetEngineOilRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EngineOil != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetEngineOilRequest](../../pkg/models/operations/getengineoilrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetEngineOilResponse](../../pkg/models/operations/getengineoilresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetFuelTank

__Description__

Returns the status of the fuel remaining in the vehicle’s gas tank. Note: The fuel tank API is only available for vehicles sold in the United States.

__Permission__

`read_fuel`

__Response Body__

|Name| Type|Desciprtion|
|--|--|--|
|`range`|number|The estimated remaining distance the car can travel (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
|`percentRemaining`|number|The remaining level of fuel in the tank (in percent).|
|`amountRemaining`|number|The amount of fuel in the tank (in liters by default or in gallons (U.S.) using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.GetFuelTank(ctx, operations.GetFuelTankRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FuelTank != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetFuelTankRequest](../../pkg/models/operations/getfueltankrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetFuelTankResponse](../../pkg/models/operations/getfueltankresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetLocation

__Description__

Retrieve latitude and longitude of a vehicle.

__Permission__

`read_location`

__Response Body__

|Name| Type|Desciprtion|
|--|--|--|
|`latitude`|number|The latitude (in degrees).|
|`longitude`|number|The longitude (in degrees).|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.GetLocation(ctx, operations.GetLocationRequest{
        VehicleID: "36ab27d0-fd9d-4455-823a-ce30af709ffc",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Location != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLocationRequest](../../pkg/models/operations/getlocationrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetLocationResponse](../../pkg/models/operations/getlocationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetOdometer

__Description__

Returns the vehicle’s last known odometer reading.

__Permission__

`read_odometer`

__Response Body__

|Name| Type|Desciprtion|
|--|--|--|
|`distance`|number|The current odometer of the vehicle (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.GetOdometer(ctx, operations.GetOdometerRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Odometer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetOdometerRequest](../../pkg/models/operations/getodometerrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetOdometerResponse](../../pkg/models/operations/getodometerresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPermissions

__Description__

Returns a list of the permissions that have been granted to your application in relation to this vehicle.

__Query Parameters__

|Parameter| Type|Required|Description|
|--|--|--|--|
|`limit`|integer|false|Number of permissions to return (default: 25).|
|`offset`|integer|false|A vehicle ID (UUID v4).|Index to start permission list at|

__Response Body__

|Name| Type|Desciprtion|
|--|--|--|
|`permissions`|array|An array of permissions.|
|`permissions.[]`|string|A permission.|
|`paging`|object|Metadata about the current list of elements.|
|`paging.count`|integer|The total number of elements for the entire query (not just the given page).|
|`paging.offset`|integer|The current start index of the returned list of elements.|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.GetPermissions(ctx, operations.GetPermissionsRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Permission != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPermissionsRequest](../../pkg/models/operations/getpermissionsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetPermissionsResponse](../../pkg/models/operations/getpermissionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetTirePressure

__Description__

Returns the air pressure of each of the vehicle’s tires.
__Permission__

`read_tires`

__Example Response__

|Name| Type|Description|
|--|--|--|
|`frontLeft`|number|The current air pressure of the front left tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
|`frontRight`|number|The current air pressure of the front right tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
|`backLeft`|number|The current air pressure of the back left tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
|`backRight`|number|The current air pressure of the back right tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.GetTirePressure(ctx, operations.GetTirePressureRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TirePressure != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTirePressureRequest](../../pkg/models/operations/gettirepressurerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetTirePressureResponse](../../pkg/models/operations/gettirepressureresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetVin

__Description__

Returns the vehicle’s manufacturer identifier.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.GetVin(ctx, operations.GetVinRequest{
        VehicleID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VinInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetVinRequest](../../pkg/models/operations/getvinrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetVinResponse](../../pkg/models/operations/getvinresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListVehicles

__Description__

Returns a paged list of all vehicles connected to the application for the current authorized user.

__Query Parameters__

|Parameter| Type|Required|Description|
|--|--|--|--|
|`limit`|integer|false|Number of vehicles to return (default: 10).|
|`offset`|integer|false|A vehicle ID (UUID v4).|Index to start vehicle list at|

__Response Body__

|Name| Type|Desciprtion|
|--|--|--|
|`vehicles`|array|An array of vehicle IDs.|
|`vehicles.[]`|string|A vehicle ID (UUID v4).|
|`paging`|object|Metadata about the current list of elements.|
|`paging.count`|integer|The total number of elements for the entire query (not just the given page).|
|`paging.offset`|integer|The current start index of the returned list of elements.|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.ListVehicles(ctx, operations.ListVehiclesRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.VehiclesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListVehiclesRequest](../../pkg/models/operations/listvehiclesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListVehiclesResponse](../../pkg/models/operations/listvehiclesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## LockUnlock

__Description__

Unlock the vehicle

__Permission__

`control_security`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.LockUnlock(ctx, operations.LockUnlockRequest{
        SecurityAction: &shared.SecurityAction{
            Action: shared.SecurityActionActionUnlock.ToPointer(),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.LockUnlockRequest](../../pkg/models/operations/lockunlockrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.LockUnlockResponse](../../pkg/models/operations/lockunlockresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## SendDestination

__Description__

Send destination coordinates to the vehicle's navigation system.

### Example Usage

```go
package main

import(
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"context"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
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
    res, err := s.Vehicles.SendDestination(ctx, operations.SendDestinationRequest{
        SendDestination: &shared.SendDestination{
            Latitude: 37.4292,
            Longitude: 122.1381,
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
| `request`                                                                                  | [operations.SendDestinationRequest](../../pkg/models/operations/senddestinationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SendDestinationResponse](../../pkg/models/operations/senddestinationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
