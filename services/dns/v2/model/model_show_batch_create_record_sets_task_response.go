package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchCreateRecordSetsTaskResponse Response Object
type ShowBatchCreateRecordSetsTaskResponse struct {

	// **参数解释：** 批量创建记录集任务的ID。 **取值范围：** 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释：** 任务状态。 **取值范围：** - PENDING：处理中 - DONE：已完成
	Status *string `json:"status,omitempty"`

	// **参数解释：** 任务的创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 任务的更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 批量创建记录集的总数。 **取值范围：** 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释：** 记录集创建成功的数量。 **取值范围：** 不涉及。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// **参数解释：** 记录集创建失败的数量。 **取值范围：** 不涉及。
	ErrorCount *int32 `json:"error_count,omitempty"`

	// **参数解释：** 记录集创建失败的条目。 **取值范围：** 不涉及。
	ErrorItems     *[]ShowBatchCreateRecordSetsTaskErrorItem `json:"error_items,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowBatchCreateRecordSetsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCreateRecordSetsTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchCreateRecordSetsTaskResponse", string(data)}, " ")
}
