package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogInfo struct {

	// **参数解释：** 检查日志页面所显示的子任务阶段名称。 **取值范围:** 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释：** 扫描日志内容。 **取值范围:** 不涉及。
	Log *string `json:"log,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围:** 不涉及。
	Level *string `json:"level,omitempty"`

	// **参数解释：** 任务失败的简要原因。 **取值范围:** 不涉及。
	Analysis *string `json:"analysis,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围:** 不涉及。
	Faq *string `json:"faq,omitempty"`
}

func (o LogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogInfo struct{}"
	}

	return strings.Join([]string{"LogInfo", string(data)}, " ")
}
