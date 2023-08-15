package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadbalancerTagsRequestBody This is a auto create Body Object
type CreateLoadbalancerTagsRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateLoadbalancerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsRequestBody", string(data)}, " ")
}
