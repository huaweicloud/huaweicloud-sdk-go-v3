package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListServerGroupsRequest struct {
	Limit  *int32  `json:"limit,omitempty"`
	Marker *string `json:"marker,omitempty"`
}

func (o ListServerGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListServerGroupsRequest", string(data)}, " ")
}
