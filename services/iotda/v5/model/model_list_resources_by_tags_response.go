package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourcesByTagsResponse struct {
	// 资源列表。
	Resources      *[]ResourceDto `json:"resources,omitempty"`
	Page           *Page          `json:"page,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListResourcesByTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsResponse", string(data)}, " ")
}
