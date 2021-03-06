package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListClusterTagsResponse struct {
	// 标签列表

	Tags           *[]TagPlain `json:"tags,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListClusterTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterTagsResponse", string(data)}, " ")
}
