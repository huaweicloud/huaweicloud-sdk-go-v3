package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisassociateRouterResponse struct {
	// Router(VPC)的ID。
	RouterId *string `json:"router_id,omitempty"`
	// Router(VPC)所在的region。
	RouterRegion *string `json:"router_region,omitempty"`
	// 资源状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateRouterResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateRouterResponse struct{}"
	}

	return strings.Join([]string{"DisassociateRouterResponse", string(data)}, " ")
}
