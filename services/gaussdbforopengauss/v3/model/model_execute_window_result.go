package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteWindowResult **参数解释**: 运维窗口。
type ExecuteWindowResult struct {

	// **参数解释**: 计划执行时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	PlannedExecutionTime string `json:"planned_execution_time"`

	// **参数解释**: 执行时间窗口开始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 执行时间窗口结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **取值范围**: 不涉及。
	EndTime string `json:"end_time"`
}

func (o ExecuteWindowResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteWindowResult struct{}"
	}

	return strings.Join([]string{"ExecuteWindowResult", string(data)}, " ")
}
