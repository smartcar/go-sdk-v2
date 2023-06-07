# Chevrolet

### Available Operations

* [GetChargeTime](#getchargetime) - Retrieve charging completion time for a Chevrolet.
* [GetVoltage](#getvoltage) - Retrieve charging voltmeter time for a Chevrolet.

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
    res, err := s.Chevrolet.GetChargeTime(ctx, operations.GetChevroletChargeTimeRequest{
        VehicleID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeTime != nil {
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
    res, err := s.Chevrolet.GetVoltage(ctx, operations.GetChevroletVoltageRequest{
        VehicleID: "quibusdam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeVoltage != nil {
        // handle response
    }
}
```
