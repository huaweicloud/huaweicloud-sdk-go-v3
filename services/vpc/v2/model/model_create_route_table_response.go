package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateRouteTableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"CreateRouteTableResponse", string(data)}, " ")
}
