# Vehicles

## Overview

Operations about vehicles

### Available Operations

* [Batch](#batch) - Batch
* [Disconnect](#disconnect) - Revoke Access
* [Get](#get) - Vehicle Info
* [GetEngineOil](#getengineoil) - Engine Oil Life
* [GetFuelTank](#getfueltank) - Fuel Tank (US Only)
* [GetLocation](#getlocation) - Location
* [GetOdometer](#getodometer) - Odometer
* [GetPermissions](#getpermissions) - Application Permissions
* [GetTirePressure](#gettirepressure) - Tire pressure
* [GetVin](#getvin) - Returns the vehicle’s manufacturer identifier.
* [ListVehicles](#listvehicles) - All Vehicles
* [LockUnlock](#lockunlock) - Lock/Unlock Vehicle

## Batch

__Description__ Returns a list of responses from multiple Smartcar endpoints, all combined into a single request. Note: Batch requests is a paid feature. Please contact us to upgrade your plan and obtain access.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.Batch(ctx, operations.BatchRequest{
        RequestBody: []string{
            "/odometer",
            "/odometer",
        },
        VehicleID: "corporis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchResponse != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.Disconnect(ctx, operations.DisconnectRequest{
        VehicleID: "iste",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Status != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.Get(ctx, operations.GetVehicleRequest{
        VehicleID: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VehicleInfo != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.GetEngineOil(ctx, operations.GetEngineOilRequest{
        VehicleID: "saepe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EngineOil != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.GetFuelTank(ctx, operations.GetFuelTankRequest{
        VehicleID: "quidem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FuelTank != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.GetOdometer(ctx, operations.GetOdometerRequest{
        VehicleID: "architecto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Odometer != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.GetPermissions(ctx, operations.GetPermissionsRequest{
        Limit: smartcarapi.Int64(60225),
        Offset: smartcarapi.Int64(969810),
        VehicleID: "est",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Permission != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.GetTirePressure(ctx, operations.GetTirePressureRequest{
        VehicleID: "mollitia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TirePressure != nil {
        // handle response
    }
}
```

## GetVin

__Description__

Returns the vehicle’s manufacturer identifier.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.GetVin(ctx, operations.GetVinRequest{
        VehicleID: "laborum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VinInfo != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.ListVehicles(ctx, operations.ListVehiclesRequest{
        Limit: smartcarapi.Int64(170909),
        Offset: smartcarapi.Int64(210382),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VehiclesResponse != nil {
        // handle response
    }
}
```

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
	"context"
	"log"
	"Smartcar-API"
	"Smartcar-API/pkg/models/operations"
	"Smartcar-API/pkg/models/shared"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.LockUnlock(ctx, operations.LockUnlockRequest{
        SecurityAction: &shared.SecurityAction{
            Action: shared.SecurityActionActionUnlock.ToPointer(),
        },
        VehicleID: "corporis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```
