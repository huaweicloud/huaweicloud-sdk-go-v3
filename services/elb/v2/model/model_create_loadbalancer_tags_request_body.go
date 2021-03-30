package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateLoadbalancerTagsRequestBody struct {
	Tag *ResourceTag `json:"tag,omitempty"`
}

func (o CreateLoadbalancerTagsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsRequestBody", string(data)}, " ")
}
