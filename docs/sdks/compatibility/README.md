# Compatibility

## Overview

Operations about compatibility

### Available Operations

* [ListCompatibility](#listcompatibility) - Compatibility

## ListCompatibility

In the US, compatibility will return a breakdown by scope of what a car is capable of. In Europe, the check is based on the make of the car so will return only a `true` or `false`

The Compatibility API lets developers determine whether a given vehicle is compatible with Smartcar and whether it is capable of the features their application requires. Using this endpoint, developers can determine whether specific vehicles are eligible before sending a user through Smartcar Connect.

A compatible vehicle is one that:
1. Has the hardware required for internet connectivity
2. Belongs to the makes and models Smartcar is compatible with in the vehicle's country

A vehicle is capable of given feature if:
1. The vehicle supports the feature (e.g., a Ford Escape supports /fuel but a Mustang Mach-e does not)
2. Smartcar supports the feature for the vehicle's make

Note: The Compatibility API is currently available for vehicles sold in the United States. For other countries, please contact us to request early access. Our initial release for vehicles sold outside of the United States only supports checking compatibility. It does not yet support checking `capabilities`.

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  compatible|   boolean|  `true` if the vehicle is likely compatible.*, `false` otherwise.| 
reason|   string \| null|  One of the enum values described below if compatible is `false`, `null` otherwise.|
| capabilities|   array| An array containing the set of endpoints that the provided scope value can provide authorization for. This array will be empty if `compatible` is false|
| capabilities[].permission|   string|One of the permissions provided in the scope parameter.|
| capabilities[].endpoint|   string| One of the endpoints that the permission authorizes access to.|
| capabilities[].capable|   boolean|`true` if the vehicle is likely capable of this feature.*, `false` otherwise.|
| capabilities[].reason|   string \| null|One of the enum values described below if capable is `false`, `null` otherwise.|

__Note__: The compatibility of a vehicle depends on many factors including its make, model, model year, trim, package, and head unit. The Smartcar Compatibility API is optimized to return false positives rather than false negatives.

__Enum Values__

|  Parameter 	|Value   	|Description   	|
|---	|---	|---	|
|  reason|   VEHICLE_NOT_COMPATIBLE|  The vehicle does not have the hardware required for internet connectivity.|
|  |   MAKE_NOT_COMPATIBLE|  Smartcar is not yet compatible with the vehicle's make in the specified country.|
|  capabilities[].reason|   VEHICLE_NOT_CAPABLE|  TThe vehicle does not support this feature.|
|  |   SMARTCAR_NOT_CAPABLE|  Smartcar is not capable of supporting the given feature on the vehicle's make.|

### Example Usage

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Compatibility.ListCompatibility(ctx, "{country}", "{scope}", "{vin}")
    if err != nil {
        log.Fatal(err)
    }

    if res.CompatibilityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `country`                                             | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | {country}                                             |
| `scope`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | {scope}                                               |
| `vin`                                                 | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   | {vin}                                                 |


### Response

**[*operations.ListCompatibilityResponse](../../models/operations/listcompatibilityresponse.md), error**

