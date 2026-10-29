package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleConfig struct {

	// **参数解释:** 规则配置ID。可调用[[获取规则列表接口](https://support.huaweicloud.com/api-codecheck/ListRules.html)](tag:hws)[[获取规则列表接口](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRules.html)](tag:hws_hk)[获取规则列表接口](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取规则配置ID。 **约束限制:** 不涉及。 **取值范围:** 32位UUID。 **默认取值:** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释:** 工具规则ID，即可配置阈值规则的各种阈值对应的ID，当前无便利手段获取，请联系客服咨询。 **约束限制:** 不涉及。 **取值范围:** 32位UUID。 **默认取值:** 不涉及。
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释:** 配置参数默认阈值。 **约束限制:** 不涉及。 **取值范围:** 不涉及。 **默认取值:** 不涉及。
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释:** 配置参数新阈值。 **约束限制:** 不涉及。 **取值范围:** 不涉及。 **默认取值:** 不涉及。
	OptionValue *string `json:"option_value,omitempty"`

	// **参数解释:** 当前规则配置项key。可调用[[获取规则列表接口](https://support.huaweicloud.com/api-codecheck/ListRules.html)](tag:hws)[[获取规则列表接口](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRules.html)](tag:hws_hk)[获取规则列表接口](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取规则配置ID。 **约束限制:** 需要和规则ID相匹配。 **取值范围:** 不涉及。 **默认取值:** 不涉及。
	OptionKey *string `json:"option_key,omitempty"`

	// **参数解释:** 当前规则配置项key。可调用[[获取规则列表接口](https://support.huaweicloud.com/api-codecheck/ListRules.html)](tag:hws)[[获取规则列表接口](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRules.html)](tag:hws_hk)[获取规则列表接口](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取规则配置ID。 **约束限制:** 需要和规则ID相匹配。 **取值范围:** 不涉及。 **默认取值:** 不涉及。
	OptionName *string `json:"option_name,omitempty"`

	// **参数解释:** 规则集ID。可调用[[查询规则集列表](https://support.huaweicloud.com/api-codecheck/ListRulesets.html)](tag:hws)[[查询规则集列表](https://support.huaweicloud.com/intl/zh-cn/api-codecheck/ListRulesets.html)](tag:hws_hk)[查询规则集列表](tag:hcs,hcs_sm,hcs_site,hcs_sitesm,fcs)获取规则集ID。 **约束限制:** 不涉及。 **取值范围:** 32位UUID。 **默认取值:** 不涉及。
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释:** 配置参数的描述。可不填。 **约束限制:** 不涉及。 **取值范围:** 不涉及。 **默认取值:** 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o RuleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleConfig struct{}"
	}

	return strings.Join([]string{"RuleConfig", string(data)}, " ")
}
