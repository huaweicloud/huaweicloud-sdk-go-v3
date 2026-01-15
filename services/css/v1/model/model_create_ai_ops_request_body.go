package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAiOpsRequestBody struct {

	// **参数解释**： 检测类型 **约束限制**： 不涉及 **取值范围**： - full_detection  全量检测项 - unavailability_detection 集群不可用检测项 - partial_detection 全量检测项中选取其中部分检测项进行检测，具体检测项需要设置check_items  **默认取值**： 不涉及
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释**： 全量检测项中选取其中部分检测项进行检测，输入检测项列表。 **约束限制**： 当check_type为partial_detection时有效 **取值范围**： 通过智能运维ShowAiOpsDetector获取最新支持的检测项，输入检测项id字符串列表 **默认取值**： 不涉及
	CheckItems *[]string `json:"check_items,omitempty"`

	// **参数解释**： 检测报告名称，支持自定义检测名。 **约束限制**： 不涉及 **取值范围**： 4～64个字符，只能包含数字、字母、中划线和下划线，且必须以字母开头。 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 检测报告描述，支持自定义检测描述 **约束限制**： 不涉及 **取值范围**： 0~255个字符 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	Alarm *CreateAiOpsRequestBodyAlarm `json:"alarm,omitempty"`
}

func (o CreateAiOpsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiOpsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAiOpsRequestBody", string(data)}, " ")
}
