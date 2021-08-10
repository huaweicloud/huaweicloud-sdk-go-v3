package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisassociateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o DisassociateRouteTableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"DisassociateRouteTableResponse", string(data)}, " ")
}
