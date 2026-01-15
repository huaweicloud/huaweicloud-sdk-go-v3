package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiOpsSetting struct {

	// **参数解释**： 检测类型。 **约束限制**： 不涉及 **取值范围**： - full_detection：  全量检测项。 - unavailability_detection： 集群不可用检测项。 - partial_detection： 全量检测项中选取其中部分检测项进行检测，具体检测项需要设置check_items。 **默认取值**： 不涉及
	CheckType string `json:"check_type"`

	// **参数解释**： 全量检测项中选取其中部分检测项进行检测，输入检测项列表。 **约束限制**： 当check_type为partial_detection时有效 **取值范围**： 通过智能运维ShowAiOpsDetector获取最新支持的检测项，输入检测项id字符串列表 **默认取值**： 不涉及
	CheckItems *[]string `json:"check_items,omitempty"`

	// **参数解释**： 智能运维自动检测时间，格式为时间加时区，例如：\"00:00 GMT+08:00\"。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Period string `json:"period"`
}

func (o AiOpsSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiOpsSetting struct{}"
	}

	return strings.Join([]string{"AiOpsSetting", string(data)}, " ")
}
