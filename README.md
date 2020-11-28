# Tines Go API Client

Go client for the Tines API

## Usage

Get Agent:

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
}
```

Create Agent:

```
package main

import (
	"fmt"

	"github.com/trivago/tgo/tcontainer"

	"github.com/tuckner/go-tines"
)

func main() {
...

	tinesClient, _ := tines.NewClient(nil, base, userEmail, userToken)

	sourceid := make([]string, 0)
	receiveid := make([]string, 0)
	custom := tcontainer.NewMarshalMap()
	custom["options"] = map[string]string{"secret": "secretphrase", "verbs": "get,post"}

	a := tines.Agent{
		Type:          "Agents::WebhookAgent",
		Name:          "Created Agent",
		StoryID:       30,
		KeepEventsFor: 604800,
		SourceIds:     sourceid,
		ReceiverIds:   receiveid,
		Unknowns:      custom,
	}

	agent, resp, err := tinesClient.Agent.Create(&a)

	fmt.Printf("%+v, %+v, %+v", agent, resp, err)
}
```

Create Global Resource:

```
...
	gr := tines.GlobalResource{
		Name:      "NewResource",
		ValueType: "json",
		Value:  "{\"key\": \"value\"}",
	}
	
	globalresource, resp, err := tinesClient.GlobalResource.Create(&gr)
...
```

## Call a not implemented API endpoint
Not all API endpoints of the Tines API are implemented into go-tines. But you can call them anyway.

```
package main

import (
	"fmt"
	"time"

	"github.com/tuckner/go-tines"
)

func main() {
	base := "tinesurl"
	userEmail := "email"
	userToken := "apitoken"
	
	tinesClient, err := tines.NewClient(nil, base, userEmail, userToken)
	req, _ := tinesClient.NewRequest("GET", "agents", nil)

	response := new(Response)
	_, err = tinesClient.Do(req, response)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(response.Agent); i++ {
		fmt.Println(response.Agent[i])
	}
}

// Response structure
type Response struct {
	Agent []Agent `json:"agents"`
}

// Agent structure
type Agent struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	Options struct {
		Mode     string `json:"mode"`
		Lookback string `json:"lookback"`
		Path     string `json:"path"`
	} `json:"options"`
	Name               string      `json:"name"`
	Schedule           interface{} `json:"schedule"`
	EventsCount        int         `json:"events_count"`
	LastCheckAt        interface{} `json:"last_check_at"`
	LastReceiveAt      time.Time   `json:"last_receive_at"`
	LastCheckedEventID int         `json:"last_checked_event_id"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	LastWebRequestAt   interface{} `json:"last_web_request_at"`
	KeepEventsFor      int         `json:"keep_events_for"`
	LastEventAt        time.Time   `json:"last_event_at"`
	LastErrorLogAt     interface{} `json:"last_error_log_at"`
	Disabled           bool        `json:"disabled"`
	GUID               string      `json:"guid"`
	StoryID            int         `json:"story_id"`
}
