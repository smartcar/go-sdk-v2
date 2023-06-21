# Webhooks

### Available Operations

* [Subscribe](#subscribe) - Subscribe Webhook
* [Unsubscribe](#unsubscribe) - Unsubscribe Webhook

## Subscribe

__Description__

Subscribe to a webhook for a vehicle.

__Permission__

`required: webhook:read`

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
	"smartcar"
	"smartcar/pkg/models/operations"
	"smartcar/pkg/models/shared"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.Subscribe(ctx, operations.SubscribeRequest{
        WebhookInfo: &shared.WebhookInfo{
            Vehicleid: smartcar.String("dc6ea99e-57d1-4e41-b129-27e7eb58713e"),
            Webhookid: smartcar.String("9b6ae692-60cc-4b3e-89d8-71e7549cf805"),
        },
        VehicleID: "repellendus",
        WebhookID: "sapiente",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.SubscribeRequest](../../models/operations/subscriberequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.SubscribeResponse](../../models/operations/subscriberesponse.md), error**


## Unsubscribe

__Description__

Delete a webhook for a vehicle.

__Permission__

`required: webhook:read`

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
	"smartcar"
	"smartcar/pkg/models/operations"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.Unsubscribe(ctx, operations.UnsubscribeRequest{
        VehicleID: "quo",
        WebhookID: "odit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UnsubscribeRequest](../../models/operations/unsubscriberequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.UnsubscribeResponse](../../models/operations/unsubscriberesponse.md), error**

