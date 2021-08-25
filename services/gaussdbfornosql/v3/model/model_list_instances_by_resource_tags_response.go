package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListInstancesByResourceTagsResponse struct {
	// 实例列表。

	Instances *[]InstanceResult `json:"instances,omitempty"`
	// 总记录数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesByResourceTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesByResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesByResourceTagsResponse", string(data)}, " ")
}
