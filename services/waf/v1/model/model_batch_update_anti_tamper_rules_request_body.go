package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateAntiTamperRulesRequestBody struct {

	// 防护网站，查询云模式防护域名列表（ListHost）接口获取防护域名，响应体中的的hostname字段
	Hostname string `json:"hostname"`

	// 防篡改规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀
	Url string `json:"url"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 策略和规则id数组，关联防护策略与对应的规则集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds []PolicyRuleIdRequestBodyPolicyRuleIds `json:"policy_rule_ids"`
}

func (o BatchUpdateAntiTamperRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAntiTamperRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateAntiTamperRulesRequestBody", string(data)}, " ")
}
