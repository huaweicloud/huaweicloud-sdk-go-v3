package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Premium struct {

	// 是否开通独享模式
	Purchased *bool `json:"purchased,omitempty"`

	// 独享实例数量，包括elb
	Total *int32 `json:"total,omitempty"`

	// elb实例数量
	Elb *int32 `json:"elb,omitempty"`

	// 独享WAF实例数量
	Dedicated *int32 `json:"dedicated,omitempty"`
}

func (o Premium) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Premium struct{}"
	}

	return strings.Join([]string{"Premium", string(data)}, " ")
}
