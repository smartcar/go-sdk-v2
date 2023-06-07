# User

### Available Operations

* [GetInfo](#getinfo) - User Info

## GetInfo

__Description__

Returns the id of the vehicle owner who granted access to your application. This should be used as the static unique identifier for storing the access token and refresh token pair in your database. Note: A single user can own multiple vehicles, and multiple users can own the same vehicle.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Smartcar-API"
)

func main() {
    s := smartcarapi.New(
        smartcarapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UserInfo != nil {
        // handle response
    }
}
```
