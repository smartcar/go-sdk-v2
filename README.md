# Smartcar Go SDK [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/smartcar/go-sdk-v2)

## Overview

The [Smartcar API](https://smartcar.com/docs) lets you read vehicle data (location, odometer, fuel, etc.) and send commands to vehicles (lock, unlock) using HTTP requests.

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/smartcar/go-sdk-v2
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Cadillac](docs/cadillac/README.md)

* [GetChargeTime](docs/cadillac/README.md#getchargetime) - Retrieve charging completion time for a Cadillac.
* [GetVoltage](docs/cadillac/README.md#getvoltage) - Retrieve charging voltmeter time for a Cadillac.

### [Chevrolet](docs/chevrolet/README.md)

* [GetChargeTime](docs/chevrolet/README.md#getchargetime) - Retrieve charging completion time for a Chevrolet.
* [GetVoltage](docs/chevrolet/README.md#getvoltage) - Retrieve charging voltmeter time for a Chevrolet.

### [Compatibility](docs/compatibility/README.md)

* [ListCompatibility](docs/compatibility/README.md#listcompatibility) - Compatibility

### [Evs](docs/evs/README.md)

* [GetBatteryCapacity](docs/evs/README.md#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](docs/evs/README.md#getbatterylevel) - EV Battery Level
* [GetChargingLimit](docs/evs/README.md#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](docs/evs/README.md#getchargingstatus) - EV Charging Status
* [SetChargingLimit](docs/evs/README.md#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](docs/evs/README.md#startstopcharge) - Start or stop charging an electric vehicle. Please contact us to request early access.

### [Tesla](docs/tesla/README.md)

* [GetAmmeter](docs/tesla/README.md#getammeter) - Retrieve charging ammeter time for a Tesla.
* [GetChargeTime](docs/tesla/README.md#getchargetime) - Retrieve charging completion time for a Tesla.
* [GetCompass](docs/tesla/README.md#getcompass) - Retrieve compass heading for a Tesla.
* [GetExteriorTemperature](docs/tesla/README.md#getexteriortemperature) - Retrieve exterior temperature for a Tesla.
* [GetInteriorTemperature](docs/tesla/README.md#getinteriortemperature) - Retrieve interior temperature for a Tesla.
* [GetSpeedometer](docs/tesla/README.md#getspeedometer) - Retrieve speed for a Tesla.
* [GetVoltage](docs/tesla/README.md#getvoltage) - Retrieve charging voltmeter time for a Tesla.
* [GetWattmeter](docs/tesla/README.md#getwattmeter) - Retrieve charging wattmeter time for a Tesla.
* [SetAmmeter](docs/tesla/README.md#setammeter) - Set charging ammeter time for a Tesla.

### [User](docs/user/README.md)

* [GetInfo](docs/user/README.md#getinfo) - User Info

### [Vehicles](docs/vehicles/README.md)

* [Batch](docs/vehicles/README.md#batch) - Batch
* [Disconnect](docs/vehicles/README.md#disconnect) - Revoke Access
* [Get](docs/vehicles/README.md#get) - Vehicle Info
* [GetEngineOil](docs/vehicles/README.md#getengineoil) - Engine Oil Life
* [GetFuelTank](docs/vehicles/README.md#getfueltank) - Fuel Tank (US Only)
* [GetLocation](docs/vehicles/README.md#getlocation) - Location
* [GetOdometer](docs/vehicles/README.md#getodometer) - Odometer
* [GetPermissions](docs/vehicles/README.md#getpermissions) - Application Permissions
* [GetTirePressure](docs/vehicles/README.md#gettirepressure) - Tire pressure
* [GetVin](docs/vehicles/README.md#getvin) - Returns the vehicle’s manufacturer identifier.
* [ListVehicles](docs/vehicles/README.md#listvehicles) - All Vehicles
* [LockUnlock](docs/vehicles/README.md#lockunlock) - Lock/Unlock Vehicle

### [Webhooks](docs/webhooks/README.md)

* [Subscribe](docs/webhooks/README.md#subscribe) - Subscribe Webhook
* [Unsubscribe](docs/webhooks/README.md#unsubscribe) - Unsubscribe Webhook
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
