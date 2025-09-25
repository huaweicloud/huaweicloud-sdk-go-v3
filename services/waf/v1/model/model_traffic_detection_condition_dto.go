package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrafficDetectionConditionDto struct {

	// **参数解释：** Id，唯一标识当前流量检测条件。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 匹配字段（类别），标识流量筛选的字段类型（如url表示URL路径）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	MatchField *string `json:"match_field,omitempty"`

	// **参数解释：** 子字段，匹配字段的细分维度（如无则不填）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	MatchFieldIndex *string `json:"match_field_index,omitempty"`

	// **参数解释：** 逻辑运算符，标识匹配条件的逻辑关系（如contain表示包含）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LogicalOperator *string `json:"logical_operator,omitempty"`

	// **参数解释：** 匹配内容，符合筛选条件的具体值列表（如特定URL路径）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	MatchContents *[]string `json:"match_contents,omitempty"`

	// **参数解释：** 引用表Id，关联预设的匹配内容列表ID（如无则不填）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ValueListRes *string `json:"value_list_res,omitempty"`
}

func (o TrafficDetectionConditionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficDetectionConditionDto struct{}"
	}

	return strings.Join([]string{"TrafficDetectionConditionDto", string(data)}, " ")
}
