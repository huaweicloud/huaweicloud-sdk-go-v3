package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateLoadbalancerTagsRequestBody struct {
	Tag *ResourceTag `json:"tag,omitempty"`
}

func (o CreateLoadbalancerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsRequestBody", string(data)}, " ")
}
