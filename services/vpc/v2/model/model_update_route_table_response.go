package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateRouteTableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableResponse", string(data)}, " ")
}
