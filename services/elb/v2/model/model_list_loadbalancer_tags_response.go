package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListLoadbalancerTagsResponse struct {
	// 标签列表

	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListLoadbalancerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"ListLoadbalancerTagsResponse", string(data)}, " ")
}
