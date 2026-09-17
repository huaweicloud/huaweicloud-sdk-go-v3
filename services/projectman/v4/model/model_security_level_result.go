package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityLevelResult **参数解释**： 密级字段信息。仅在涉密环境（SM）下存在此字段，非涉密环境下无此字段。
type SecurityLevelResult struct {

	// **参数解释**： 密级字段ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 密级字段名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DisplayValue *string `json:"display_value,omitempty"`

	// **参数解释**： 用户自定义的密级字段的值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 密级编码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 密级排序值，越大级别越高。 **取值范围**： 不涉及。
	Sequence *float64 `json:"sequence,omitempty"`
}

func (o SecurityLevelResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityLevelResult struct{}"
	}

	return strings.Join([]string{"SecurityLevelResult", string(data)}, " ")
}
