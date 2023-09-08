# Smartcar Go SDK [![Build Status][ci-image]][ci-url] [![Release][release-image]][release-url] [![GoDoc][doc-image]][doc-url]

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
	"github.com/smartcar/go-sdk-v2"
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"github.com/smartcar/go-sdk-v2/pkg/models/operations"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Cadillac](docs/sdks/cadillac/README.md)

* [GetChargeTime](docs/sdks/cadillac/README.md#getchargetime) - Retrieve charging completion time for a Cadillac.
* [GetVoltage](docs/sdks/cadillac/README.md#getvoltage) - Retrieve charging voltmeter time for a Cadillac.

### [Chevrolet](docs/sdks/chevrolet/README.md)

* [GetChargeTime](docs/sdks/chevrolet/README.md#getchargetime) - Retrieve charging completion time for a Chevrolet.
* [GetVoltage](docs/sdks/chevrolet/README.md#getvoltage) - Retrieve charging voltmeter time for a Chevrolet.

### [Compatibility](docs/sdks/compatibility/README.md)

* [ListCompatibility](docs/sdks/compatibility/README.md#listcompatibility) - Compatibility

### [Evs](docs/sdks/evs/README.md)

* [GetBatteryCapacity](docs/sdks/evs/README.md#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](docs/sdks/evs/README.md#getbatterylevel) - EV Battery Level
* [GetChargingLimit](docs/sdks/evs/README.md#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](docs/sdks/evs/README.md#getchargingstatus) - EV Charging Status
* [SetChargingLimit](docs/sdks/evs/README.md#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](docs/sdks/evs/README.md#startstopcharge) - Start or stop charging an electric vehicle.

### [Tesla](docs/sdks/tesla/README.md)

* [GetAmmeter](docs/sdks/tesla/README.md#getammeter) - Retrieve charging ammeter time for a Tesla.
* [GetChargeTime](docs/sdks/tesla/README.md#getchargetime) - Retrieve charging completion time for a Tesla.
* [GetCompass](docs/sdks/tesla/README.md#getcompass) - Retrieve compass heading for a Tesla.
* [GetExteriorTemperature](docs/sdks/tesla/README.md#getexteriortemperature) - Retrieve exterior temperature for a Tesla.
* [GetInteriorTemperature](docs/sdks/tesla/README.md#getinteriortemperature) - Retrieve interior temperature for a Tesla.
* [GetSpeedometer](docs/sdks/tesla/README.md#getspeedometer) - Retrieve speed for a Tesla.
* [GetVoltage](docs/sdks/tesla/README.md#getvoltage) - Retrieve charging voltmeter time for a Tesla.
* [GetWattmeter](docs/sdks/tesla/README.md#getwattmeter) - Retrieve charging wattmeter time for a Tesla.
* [SetAmmeter](docs/sdks/tesla/README.md#setammeter) - Set charging ammeter time for a Tesla.

### [User](docs/sdks/user/README.md)

* [GetInfo](docs/sdks/user/README.md#getinfo) - User Info

### [VehicleManagement](docs/sdks/vehiclemanagement/README.md)

* [DeleteManagementVehicleConnections](docs/sdks/vehiclemanagement/README.md#deletemanagementvehicleconnections) - Delete vehicle connections by user_id or vehicle_id
* [GetManagementVehicleConnections](docs/sdks/vehiclemanagement/README.md#getmanagementvehicleconnections) - Retrieve vehicle connections

### [Vehicles](docs/sdks/vehicles/README.md)

* [Batch](docs/sdks/vehicles/README.md#batch) - Batch
* [LockStatus](docs/sdks/vehicles/README.md#lockstatus) - Returns the lock status for a vehicle and the open status of its doors, windows, storage units, sunroof and charging port where available. The open status array(s) will be empty if a vehicle has partial support. The request will error if lock status can not be retrieved from the vehicle or the brand is not supported.
* [Disconnect](docs/sdks/vehicles/README.md#disconnect) - Revoke Access
* [Get](docs/sdks/vehicles/README.md#get) - Vehicle Info
* [GetEngineOil](docs/sdks/vehicles/README.md#getengineoil) - Engine Oil Life
* [GetFuelTank](docs/sdks/vehicles/README.md#getfueltank) - Fuel Tank (US Only)
* [GetLocation](docs/sdks/vehicles/README.md#getlocation) - Location
* [GetOdometer](docs/sdks/vehicles/README.md#getodometer) - Odometer
* [GetPermissions](docs/sdks/vehicles/README.md#getpermissions) - Application Permissions
* [GetTirePressure](docs/sdks/vehicles/README.md#gettirepressure) - Tire Pressure
* [GetVin](docs/sdks/vehicles/README.md#getvin) - Returns the vehicle’s manufacturer identifier.
* [ListVehicles](docs/sdks/vehicles/README.md#listvehicles) - All Vehicles
* [LockUnlock](docs/sdks/vehicles/README.md#lockunlock) - Lock/Unlock Vehicle

### [Webhooks](docs/sdks/webhooks/README.md)

* [Subscribe](docs/sdks/webhooks/README.md#subscribe) - Subscribe Webhook
* [Unsubscribe](docs/sdks/webhooks/README.md#unsubscribe) - Unsubscribe Webhook
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

[ci-url]: https://github.com/smartcar/go-sdk-v2/actions/workflows/speakeasy_sdk_generation.yml
[ci-image]: https://github.com/smartcar/go-sdk-v2/actions/workflows/speakeasy_sdk_generation.yml/badge.svg
[release-url]: https://github.com/smartcar/go-sdk-v2/releases
[release-image]: https://img.shields.io/github/v/release/smartcar/go-sdk-v2?sort=semver
[doc-url]: http://godoc.org/github.com/smartcar/go-sdk-v2
[doc-image]: http://img.shields.io/badge/godoc-reference-blue.svg
