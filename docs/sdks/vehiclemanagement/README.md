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
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
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
    res, err := s.VehicleManagement.DeleteManagementVehicleConnections(ctx, operations.DeleteManagementVehicleConnectionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DeletedConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.DeleteManagementVehicleConnectionsRequest](../../pkg/models/operations/deletemanagementvehicleconnectionsrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |


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
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/shared"
	gosdkv2 "github.com/smartcar/go-sdk-v2/v3"
	"context"
	"github.com/smartcar/go-sdk-v2/v3/pkg/models/operations"
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
    res, err := s.VehicleManagement.GetManagementVehicleConnections(ctx, operations.GetManagementVehicleConnectionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetManagementVehicleConnectionsRequest](../../pkg/models/operations/getmanagementvehicleconnectionsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |


### Response

**[*operations.GetManagementVehicleConnectionsResponse](../../pkg/models/operations/getmanagementvehicleconnectionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
