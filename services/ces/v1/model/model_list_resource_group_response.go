package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourceGroupResponse struct {
	// 一个或者多个资源分组信息。

	ResourceGroups *[]ResourceGroupInfo `json:"resource_groups,omitempty"`

	MetaData       *TotalMetaData `json:"meta_data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListResourceGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ListResourceGroupResponse", string(data)}, " ")
}
