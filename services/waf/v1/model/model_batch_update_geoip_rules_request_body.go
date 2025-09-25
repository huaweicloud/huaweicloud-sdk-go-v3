package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateGeoipRulesRequestBody **参数解释：** 批量更新地理位置封禁body，用于批量修改地理位置访问控制规则的配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BatchUpdateGeoipRulesRequestBody struct {

	// **参数解释：** 策略和规则id数组，关联防护策略与对应的规则集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds *[]PolicyRuleIdRequestBodyPolicyRuleIds `json:"policy_rule_ids,omitempty"`

	// **参数解释：** 规则状态，控制地理位置访问控制规则的启用/禁用（如1表示启用，0表示禁用） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 规则名称，标识地理位置访问控制规则的名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 地理位置，指定需要控制的地域（如省份、城市编码，多个用|分隔） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Geoip *string `json:"geoip,omitempty"`

	// **参数解释：** 放行或者拦截，1表示放行指定地理位置，2表示拦截指定地理位置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	White *int32 `json:"white,omitempty"`
}

func (o BatchUpdateGeoipRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateGeoipRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateGeoipRulesRequestBody", string(data)}, " ")
}
