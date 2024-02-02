# Webhooks
(*Webhooks*)

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
    res, err := s.Webhooks.Subscribe(ctx, operations.SubscribeRequest{
        WebhookInfo: &shared.WebhookInfo{
            Vehicleid: gosdkv2.String("dc6ea99e-57d1-4e41-b129-27e7eb58713e"),
            Webhookid: gosdkv2.String("9b6ae692-60cc-4b3e-89d8-71e7549cf805"),
        },
        VehicleID: "string",
        WebhookID: "string",
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
| `request`                                                                      | [operations.SubscribeRequest](../../pkg/models/operations/subscriberequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.SubscribeResponse](../../pkg/models/operations/subscriberesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
    res, err := s.Webhooks.Unsubscribe(ctx, operations.UnsubscribeRequest{
        VehicleID: "string",
        WebhookID: "string",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UnsubscribeRequest](../../pkg/models/operations/unsubscriberequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UnsubscribeResponse](../../pkg/models/operations/unsubscriberesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
