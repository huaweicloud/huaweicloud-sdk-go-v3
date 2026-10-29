package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomAttributesRule struct {

	// **参数解释:** 规则ID。可调用[[获取规则列表接口](https://support.huaweicloud.com/api-codecheck/ListRules.html)](tag:hws)[[获取规则列表接口](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRules.html)](tag:hws_hk)[获取规则列表接口](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取。 **约束限制:** 不涉及。 **取值范围:** 32位项目ID。 **默认取值:** 不涉及。
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释:** attribute的告警等级。 **约束限制:** 不涉及。 **取值范围:** - 0：致命。 - 1：严重。 - 2：一般。 - 3：提示。 **默认取值:** 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释:** 规则阈值详细。
	RuleConfigList *[]RuleConfig `json:"rule_config_list,omitempty"`
}

func (o CustomAttributesRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomAttributesRule struct{}"
	}

	return strings.Join([]string{"CustomAttributesRule", string(data)}, " ")
}
