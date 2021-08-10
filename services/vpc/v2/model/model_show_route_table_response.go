package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowRouteTableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRouteTableResponse struct{}"
	}

	return strings.Join([]string{"ShowRouteTableResponse", string(data)}, " ")
}
