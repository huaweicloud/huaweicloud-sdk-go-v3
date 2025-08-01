package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRuleFixParamInfo 配置检测检查项参数信息
type CheckRuleFixParamInfo struct {

	// **参数解释**: 检查项参数ID **取值范围**: 不涉及
	RuleParamId *int32 `json:"rule_param_id,omitempty"`

	// **参数解释**: 检查项参数描述 **取值范围**: 不涉及
	RuleDesc *string `json:"rule_desc,omitempty"`

	// **参数解释**: 检查项参数默认值 **取值范围**: 不涉及
	DefaultValue *int32 `json:"default_value,omitempty"`

	// **参数解释**: 检查项参数可取最小值 **取值范围**: 不涉及
	RangeMin *int32 `json:"range_min,omitempty"`

	// **参数解释**: 检查项参数可取最大值 **取值范围**: 不涉及
	RangeMax *int32 `json:"range_max,omitempty"`
}

func (o CheckRuleFixParamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleFixParamInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleFixParamInfo", string(data)}, " ")
}
