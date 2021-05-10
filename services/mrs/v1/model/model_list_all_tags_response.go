package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAllTagsResponse struct {
	// 标签列表信息

	Tags           *[]TagWithMultiValue `json:"tags,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAllTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAllTagsResponse struct{}"
	}

	return strings.Join([]string{"ListAllTagsResponse", string(data)}, " ")
}
