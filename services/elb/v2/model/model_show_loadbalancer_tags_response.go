package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowLoadbalancerTagsResponse struct {
	// 标签列表

	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowLoadbalancerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerTagsResponse", string(data)}, " ")
}
