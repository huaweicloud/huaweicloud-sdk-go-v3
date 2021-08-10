package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRouteTableResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRouteTableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRouteTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteRouteTableResponse", string(data)}, " ")
}
