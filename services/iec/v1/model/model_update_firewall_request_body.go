package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFirewallRequestBody 更新网络ACL请求体
type UpdateFirewallRequestBody struct {
	Firewall *UpdateFirewallOption `json:"firewall"`
}

func (o UpdateFirewallRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRequestBody", string(data)}, " ")
}
