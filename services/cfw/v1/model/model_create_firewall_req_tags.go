package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFirewallReqTags struct {

	// 资源标签键
	Key *string `json:"key,omitempty"`

	// 资源标签值
	Value *string `json:"value,omitempty"`
}

func (o CreateFirewallReqTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallReqTags struct{}"
	}

	return strings.Join([]string{"CreateFirewallReqTags", string(data)}, " ")
}
