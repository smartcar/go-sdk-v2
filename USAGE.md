<!-- Start SDK Example Usage [usage] -->
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