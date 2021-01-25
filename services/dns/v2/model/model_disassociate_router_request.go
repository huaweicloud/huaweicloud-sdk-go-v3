package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisassociateRouterRequest struct {
	ZoneId string                 `json:"zone_id"`
	Body   *DisassociaterouterReq `json:"body,omitempty"`
}

func (o DisassociateRouterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateRouterRequest struct{}"
	}

	return strings.Join([]string{"DisassociateRouterRequest", string(data)}, " ")
}
