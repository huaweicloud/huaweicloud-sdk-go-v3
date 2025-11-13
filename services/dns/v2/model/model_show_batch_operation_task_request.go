package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchOperationTaskRequest Request Object
type ShowBatchOperationTaskRequest struct {

	// **参数解释：** 批量操作任务的ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TaskId string `json:"task_id"`

	// **参数解释：** 分页查询时配置每页返回的失败条目个数。 **约束限制：** 不涉及。 **取值范围：** 0~500。 **默认取值：** 500
	ErrorItemLimit *int32 `json:"error_item_limit,omitempty"`

	// **参数解释：** 分页查询起始偏移量，表示从偏移量的下一个失败条目开始查询。 **约束限制：** 不涉及。 **取值范围：** 0~2147483647。 **默认取值：** 0
	ErrorItemOffset *int32 `json:"error_item_offset,omitempty"`
}

func (o ShowBatchOperationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchOperationTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchOperationTaskRequest", string(data)}, " ")
}
