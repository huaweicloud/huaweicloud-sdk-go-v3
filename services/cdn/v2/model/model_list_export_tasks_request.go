package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExportTasksRequest Request Object
type ListExportTasksRequest struct {

	// **参数解释：** 每页显示的条目数量 **约束限制：** 不涉及 **取值范围：** 0-100 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 偏移量 > 表示从此偏移量开始查询  **约束限制：** 不涉及 **取值范围：** offset大于等于0 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 任务id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释：** 任务名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TaskName *string `json:"task_name,omitempty"`
}

func (o ListExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListExportTasksRequest", string(data)}, " ")
}
