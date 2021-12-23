package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建网络ACL请求体。
type CreateFirewallRequestBody struct {
	Firewall *CreateFirewallOption `json:"firewall"`
}

func (o CreateFirewallRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFirewallRequestBody", string(data)}, " ")
}
