package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListOrgInstancesResponse struct {
	Instances *PageInstancesVo `json:"instances,omitempty"`
	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListOrgInstancesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOrgInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListOrgInstancesResponse", string(data)}, " ")
}
