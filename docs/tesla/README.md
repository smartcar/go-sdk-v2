# Tesla

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
    res, err := s.Tesla.GetAmmeter(ctx, operations.GetTeslaAmmeterRequest{
        VehicleID: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeAmmeter != nil {
        // handle response
    }
}
```

## GetChargeTime

__Description__

When the vehicle is charging, this endpoint returns the date and time the vehicle expects to complete this charging session. When the vehicle is not charging, this endpoint results in a vehicle state error.

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
    res, err := s.Tesla.GetChargeTime(ctx, operations.GetTeslaChargeTimeRequest{
        VehicleID: "suscipit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeTime != nil {
        // handle response
    }
}
```

## GetCompass

__Description__

This endpoint returns the compass heading of a Tesla. The value is in degrees, with 0 degrees being North.

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
    res, err := s.Tesla.GetCompass(ctx, operations.GetTeslaCompassRequest{
        VehicleID: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Compass != nil {
        // handle response
    }
}
```

## GetExteriorTemperature

__Description__

This endpoint returns the exterior temperature of a Tesla, in celsius by default.

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
    res, err := s.Tesla.GetExteriorTemperature(ctx, operations.GetTeslaExteriorTemperatureRequest{
        ID: "4e0f467c-c879-46ed-951a-05dfc2ddf7cc",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Temperature != nil {
        // handle response
    }
}
```

## GetInteriorTemperature

__Description__

This endpoint returns the interior temperature of a Tesla, in celsius by default.

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
    res, err := s.Tesla.GetInteriorTemperature(ctx, operations.GetTeslaInteriorTemperatureRequest{
        ID: "78ca1ba9-28fc-4816-b42c-b73920592939",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Temperature != nil {
        // handle response
    }
}
```

## GetSpeedometer

__Description__

This endpoint returns the speed of a Tesla (in kilometers/hour by default or in miles/hour using the sc-unit-system).

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
    res, err := s.Tesla.GetSpeedometer(ctx, operations.GetTeslaSpeedRequest{
        VehicleID: "laboriosam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Speed != nil {
        // handle response
    }
}
```

## GetVoltage

__Description__

When the vehicle is plugged in, this endpoint returns the voltage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

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
    res, err := s.Tesla.GetVoltage(ctx, operations.GetTeslaVoltageRequest{
        VehicleID: "hic",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeVoltage != nil {
        // handle response
    }
}
```

## GetWattmeter

__Description__

When the vehicle is plugged in, this endpoint returns the wattage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

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
    res, err := s.Tesla.GetWattmeter(ctx, operations.GetTeslaWattmeterRequest{
        VehicleID: "saepe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeWattage != nil {
        // handle response
    }
}
```

## SetAmmeter

__Description__

When the vehicle is plugged in, this endpoint sets the amperage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

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
    res, err := s.Tesla.SetAmmeter(ctx, operations.SetTeslaAmmeterRequest{
        ChargeAmmeter: &shared.ChargeAmmeter{
            Amperage: smartcarapi.Float32(48),
        },
        VehicleID: "fuga",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```
