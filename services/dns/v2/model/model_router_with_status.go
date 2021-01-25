package model

import (
	"encoding/json"

	"strings"
)

type RouterWithStatus struct {
	// 资源状态。
	Status *string `json:"status,omitempty"`
	// Router(VPC)所属VPC的ID。
	RouterId *string `json:"router_id,omitempty"`
	// Router(VPC)所在的region。
	RouterRegion *string `json:"router_region,omitempty"`
}

func (o RouterWithStatus) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RouterWithStatus struct{}"
	}

	return strings.Join([]string{"RouterWithStatus", string(data)}, " ")
}
