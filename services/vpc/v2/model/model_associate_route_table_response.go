package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o AssociateRouteTableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"AssociateRouteTableResponse", string(data)}, " ")
}
