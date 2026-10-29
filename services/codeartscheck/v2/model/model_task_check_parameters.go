package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskCheckParameters **参数解释：** 任务规则集的检查参数。
type TaskCheckParameters struct {

	// **参数解释：** 检查工具ID。 **取值范围:** ≥0
	CheckId *int32 `json:"check_id,omitempty"`

	// **参数解释：** 编译参数名称。 **取值范围:** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 检查参数配置信息。
	CheckerConfigs *[]CheckConfigInfo `json:"checker_configs,omitempty"`
}

func (o TaskCheckParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskCheckParameters struct{}"
	}

	return strings.Join([]string{"TaskCheckParameters", string(data)}, " ")
}
