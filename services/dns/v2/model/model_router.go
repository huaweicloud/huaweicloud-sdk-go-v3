package model

import (
	"encoding/json"

	"strings"
)

type Router struct {
	// Router(VPC)所属VPC的ID。
	RouterId string `json:"router_id"`
	// Router(VPC)所在的region。
	RouterRegion *string `json:"router_region,omitempty"`
}

func (o Router) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Router struct{}"
	}

	return strings.Join([]string{"Router", string(data)}, " ")
}
