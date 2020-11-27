# Tines Go API Client

Go API client for the Tines API

## Usage

```
package main

import (
	"fmt"

	"github.com/tuckner/go-tines"
)

func main() {
	base := "tinesurl"
	userEmail := "email"
	userToken := "apitoken"
  
	tinesClient, err := tines.NewClient(nil, base, userEmail, userToken)
	agent, resp, err := tinesClient.Agent.Get(173)
	fmt.Printf("%+v, %+v, %+v", agent, resp, err)
```
