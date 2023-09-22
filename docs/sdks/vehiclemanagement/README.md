# VehicleManagement

### Available Operations

* [DeleteManagementVehicleConnections](#deletemanagementvehicleconnections) - Delete vehicle connections by user_id or vehicle_id
* [GetManagementVehicleConnections](#getmanagementvehicleconnections) - Retrieve vehicle connections

## DeleteManagementVehicleConnections

Delete all connections by vehicle or user ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkv2 "github.com/smartcar/go-sdk-v2"
	"github.com/smartcar/go-sdk-v2/pkg/models/operations"
)

func main() {
    s := gosdkv2.New()
    operationSecurity := operations.DeleteManagementVehicleConnectionsSecurity{
            Password: "",
            Username: "",
        }

    ctx := context.Background()
    res, err := s.VehicleManagement.DeleteManagementVehicleConnections(ctx, operations.DeleteManagementVehicleConnectionsRequest{
        UserID: gosdkv2.String("in"),
        VehicleID: gosdkv2.String("corporis"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletedConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteManagementVehicleConnectionsRequest](../../models/operations/deletemanagementvehicleconnectionsrequest.md)   | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `security`                                                                                                                     | [operations.DeleteManagementVehicleConnectionsSecurity](../../models/operations/deletemanagementvehicleconnectionssecurity.md) | :heavy_check_mark:                                                                                                             | The security requirements to use for the request.                                                                              |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |


### Response

**[*operations.DeleteManagementVehicleConnectionsResponse](../../models/operations/deletemanagementvehicleconnectionsresponse.md), error**


## GetManagementVehicleConnections

Returns a paged list of all the vehicles that are connected to the application associated with the management API token used sorted in descending order by connection date.

### Example Usage

```go
package main

import(
	"context"
	"log"
	gosdkv2 "github.com/smartcar/go-sdk-v2"
	"github.com/smartcar/go-sdk-v2/pkg/models/operations"
)

func main() {
    s := gosdkv2.New()
    operationSecurity := operations.GetManagementVehicleConnectionsSecurity{
            Password: "",
            Username: "",
        }

    ctx := context.Background()
    res, err := s.VehicleManagement.GetManagementVehicleConnections(ctx, operations.GetManagementVehicleConnectionsRequest{
        Cursor: gosdkv2.String("iste"),
        Limit: gosdkv2.Int64(437032),
        UserID: gosdkv2.String("saepe"),
        VehicleID: gosdkv2.String("quidem"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetManagementVehicleConnectionsRequest](../../models/operations/getmanagementvehicleconnectionsrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.GetManagementVehicleConnectionsSecurity](../../models/operations/getmanagementvehicleconnectionssecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |


### Response

**[*operations.GetManagementVehicleConnectionsResponse](../../models/operations/getmanagementvehicleconnectionsresponse.md), error**

