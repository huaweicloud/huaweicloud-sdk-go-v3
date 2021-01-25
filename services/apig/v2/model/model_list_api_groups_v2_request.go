package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApiGroupsV2Request struct {
	ProjectId     string  `json:"project_id"`
	InstanceId    string  `json:"instance_id"`
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	Offset        *int64  `json:"offset,omitempty"`
	Limit         *int32  `json:"limit,omitempty"`
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListApiGroupsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiGroupsV2Request struct{}"
	}

	return strings.Join([]string{"ListApiGroupsV2Request", string(data)}, " ")
}
