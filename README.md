# Smartcar Go SDK [![Build Status][ci-image]][ci-url] [![Release][release-image]][release-url] [![GoDoc][doc-image]][doc-url]

## Overview

The [Smartcar API](https://smartcar.com/docs) lets you read vehicle data (location, odometer, fuel, etc.) and send commands to vehicles (lock, unlock) using HTTP requests.

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/smartcar/go-sdk-v2
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Compatibility](docs/sdks/compatibility/README.md)

* [ListCompatibility](docs/sdks/compatibility/README.md#listcompatibility) - Compatibility

### [VehicleManagement](docs/sdks/vehiclemanagement/README.md)

* [DeleteManagementVehicleConnections](docs/sdks/vehiclemanagement/README.md#deletemanagementvehicleconnections) - Delete vehicle connections by user_id or vehicle_id
* [GetManagementVehicleConnections](docs/sdks/vehiclemanagement/README.md#getmanagementvehicleconnections) - Retrieve vehicle connections

### [User](docs/sdks/user/README.md)

* [GetInfo](docs/sdks/user/README.md#getinfo) - User Info

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
* [SendDestination](docs/sdks/vehicles/README.md#senddestination) - Send Destination to vehicle navigation system

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

### [Evs](docs/sdks/evs/README.md)

* [GetBatteryCapacity](docs/sdks/evs/README.md#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](docs/sdks/evs/README.md#getbatterylevel) - EV Battery Level
* [GetChargingLimit](docs/sdks/evs/README.md#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](docs/sdks/evs/README.md#getchargingstatus) - EV Charging Status
* [SetChargingLimit](docs/sdks/evs/README.md#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](docs/sdks/evs/README.md#startstopcharge) - Start or stop charging an electric vehicle.

### [Cadillac](docs/sdks/cadillac/README.md)

* [GetChargeTime](docs/sdks/cadillac/README.md#getchargetime) - Retrieve charging completion time for a Cadillac.
* [GetVoltage](docs/sdks/cadillac/README.md#getvoltage) - Retrieve charging voltmeter time for a Cadillac.

### [Chevrolet](docs/sdks/chevrolet/README.md)

* [GetChargeTime](docs/sdks/chevrolet/README.md#getchargetime) - Retrieve charging completion time for a Chevrolet.
* [GetVoltage](docs/sdks/chevrolet/README.md#getvoltage) - Retrieve charging voltmeter time for a Chevrolet.

### [Webhooks](docs/sdks/webhooks/README.md)

* [Subscribe](docs/sdks/webhooks/README.md#subscribe) - Subscribe Webhook
* [Unsubscribe](docs/sdks/webhooks/README.md#unsubscribe) - Unsubscribe Webhook
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/sdkerrors"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
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
	res, err := s.Compatibility.ListCompatibility(ctx, operations.ListCompatibilityRequest{
		Country: gosdkv2.String("{country}"),
		Scope:   gosdkv2.String("{scope}"),
		Vin:     gosdkv2.String("{vin}"),
	})
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.smartcar.com/v2.0` | None |

#### Example

```go
package main

import (
	"context"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	"log"
)

func main() {
	s := gosdkv2.New(
		gosdkv2.WithServerIndex(0),
		gosdkv2.WithSecurity(shared.Security{
			BasicAuth: &shared.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Compatibility.ListCompatibility(ctx, operations.ListCompatibilityRequest{
		Country: gosdkv2.String("{country}"),
		Scope:   gosdkv2.String("{scope}"),
		Vin:     gosdkv2.String("{vin}"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CompatibilityResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	"log"
)

func main() {
	s := gosdkv2.New(
		gosdkv2.WithServerURL("https://api.smartcar.com/v2.0"),
		gosdkv2.WithSecurity(shared.Security{
			BasicAuth: &shared.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Compatibility.ListCompatibility(ctx, operations.ListCompatibilityRequest{
		Country: gosdkv2.String("{country}"),
		Scope:   gosdkv2.String("{scope}"),
		Vin:     gosdkv2.String("{vin}"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CompatibilityResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
	s := gosdkv2.New()

	operationSecurity := operations.DeleteManagementVehicleConnectionsSecurity{
		Password: "<YOUR_PASSWORD_HERE>",
		Username: "<YOUR_USERNAME_HERE>",
	}

	ctx := context.Background()
	res, err := s.VehicleManagement.DeleteManagementVehicleConnections(ctx, operations.DeleteManagementVehicleConnectionsRequest{}, operationSecurity, operations.WithServerURL("https://management.smartcar.com/v2.0"))
	if err != nil {
		log.Fatal(err)
	}

	if res.DeletedConnectionsResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BasicAuth`  | http         | HTTP Basic   |
| `BearerAuth` | http         | HTTP Bearer  |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
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
	res, err := s.Compatibility.ListCompatibility(ctx, operations.ListCompatibilityRequest{
		Country: gosdkv2.String("{country}"),
		Scope:   gosdkv2.String("{scope}"),
		Vin:     gosdkv2.String("{vin}"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CompatibilityResponse != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
	"log"
)

func main() {
	s := gosdkv2.New()

	operationSecurity := operations.DeleteManagementVehicleConnectionsSecurity{
		Password: "<YOUR_PASSWORD_HERE>",
		Username: "<YOUR_USERNAME_HERE>",
	}

	ctx := context.Background()
	res, err := s.VehicleManagement.DeleteManagementVehicleConnections(ctx, operations.DeleteManagementVehicleConnectionsRequest{}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.DeletedConnectionsResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



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
