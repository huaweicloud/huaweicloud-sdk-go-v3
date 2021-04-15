package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSubnetsByTagsResponse struct {
	// 资源列表

	Resources *[]ListResourceResp `json:"resources,omitempty"`
	// 资源数量

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubnetsByTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubnetsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsByTagsResponse", string(data)}, " ")
}
