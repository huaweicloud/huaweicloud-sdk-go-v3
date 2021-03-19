package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateRouterRequest struct {
	ZoneId string `json:"zone_id"`

	Body *AssociateRouterReq `json:"body,omitempty"`
}

func (o AssociateRouterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateRouterRequest struct{}"
	}

	return strings.Join([]string{"AssociateRouterRequest", string(data)}, " ")
}
