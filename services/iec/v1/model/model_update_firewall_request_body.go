package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新网络ACL请求体
type UpdateFirewallRequestBody struct {
	Firewall *UpdateFirewallOption `json:"firewall" xml:"firewall"`
}

func (o UpdateFirewallRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRequestBody", string(data)}, " ")
}
