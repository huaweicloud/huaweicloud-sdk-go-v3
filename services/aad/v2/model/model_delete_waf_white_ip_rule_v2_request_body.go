package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteWafWhiteIpRuleV2RequestBody struct {

	// 待删除规则id
	Ids []string `json:"ids"`

	// 域名+端口组合，标准端口80/443无须加端口。
	DomainName string `json:"domain_name"`

	// 防护区域,0-大陆,1-海外
	OverseasType int32 `json:"overseas_type"`
}

func (o DeleteWafWhiteIpRuleV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWafWhiteIpRuleV2RequestBody struct{}"
	}

	return strings.Join([]string{"DeleteWafWhiteIpRuleV2RequestBody", string(data)}, " ")
}
