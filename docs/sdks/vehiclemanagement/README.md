# VehicleManagement
(*VehicleManagement*)

## Overview

Operations to manage vehicle connections

### Available Operations

* [DeleteManagementVehicleConnections](#deletemanagementvehicleconnections) - Delete vehicle connections by user_id or vehicle_id
* [GetManagementVehicleConnections](#getmanagementvehicleconnections) - Retrieve vehicle connections

## DeleteManagementVehicleConnections

Delete all connections by vehicle or user ID.

### Example Usage

```go
package main

import(
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
	"context"
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

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.DeleteManagementVehicleConnectionsRequest](../../pkg/models/operations/deletemanagementvehicleconnectionsrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.DeleteManagementVehicleConnectionsSecurity](../../pkg/models/operations/deletemanagementvehicleconnectionssecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |


### Response

**[*operations.DeleteManagementVehicleConnectionsResponse](../../pkg/models/operations/deletemanagementvehicleconnectionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetManagementVehicleConnections

Returns a paged list of all the vehicles that are connected to the application associated with the management API token used sorted in descending order by connection date.

### Example Usage

```go
package main

import(
	gosdkv2 "github.com/smartcar/go-sdk-v2/v2"
	"github.com/smartcar/go-sdk-v2/v2/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := gosdkv2.New()


    operationSecurity := operations.GetManagementVehicleConnectionsSecurity{
            Password: "<YOUR_PASSWORD_HERE>",
            Username: "<YOUR_USERNAME_HERE>",
        }

    ctx := context.Background()
    res, err := s.VehicleManagement.GetManagementVehicleConnections(ctx, operations.GetManagementVehicleConnectionsRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetManagementVehicleConnectionsRequest](../../pkg/models/operations/getmanagementvehicleconnectionsrequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.GetManagementVehicleConnectionsSecurity](../../pkg/models/operations/getmanagementvehicleconnectionssecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |


### Response

**[*operations.GetManagementVehicleConnectionsResponse](../../pkg/models/operations/getmanagementvehicleconnectionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
