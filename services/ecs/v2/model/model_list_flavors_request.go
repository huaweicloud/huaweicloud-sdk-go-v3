package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
	// 可用区，需要指定可用区（AZ）的名称或者ID或者code。

	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
