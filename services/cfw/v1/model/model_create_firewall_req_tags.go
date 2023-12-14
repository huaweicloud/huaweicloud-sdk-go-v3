package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFirewallReqTags struct {

	// 标签KEY
	Key string `json:"key"`

	// 标签VALUE
	Value string `json:"value"`
}

func (o CreateFirewallReqTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallReqTags struct{}"
	}

	return strings.Join([]string{"CreateFirewallReqTags", string(data)}, " ")
}
