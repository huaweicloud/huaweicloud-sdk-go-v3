package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddWafWhiteIpRuleV2RequestBody struct {

	// 域名+端口组合，标准端口80/443无须加端口。
	DomainName string `json:"domain_name"`

	// 待添加ip/ip段
	Ips []string `json:"ips"`

	// 防护区域,0-大陆,1-海外
	OverseasType int32 `json:"overseas_type"`

	// 0-黑名单，1-白名单
	Type int32 `json:"type"`
}

func (o AddWafWhiteIpRuleV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWafWhiteIpRuleV2RequestBody struct{}"
	}

	return strings.Join([]string{"AddWafWhiteIpRuleV2RequestBody", string(data)}, " ")
}
