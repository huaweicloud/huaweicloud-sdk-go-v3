package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchOperationTaskResponse Response Object
type ShowBatchOperationTaskResponse struct {

	// **参数解释：** 批量操作任务的ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 任务类型。 **取值范围：** - batch_create_zone：批量添加域名 - create：批量添加记录集 - import_recordset：批量导入记录集 - delete：批量删除记录集 - batch_update_rs：批量修改记录集 - transfer：批量转移域名
	Type *string `json:"type,omitempty"`

	// **参数解释：** 任务状态。 **取值范围：** - PENDING：执行中 - DONE：已完成
	Status *string `json:"status,omitempty"`

	// **参数解释：** 任务的创建时间。 格式：yyyy-MM-dd HH:mm:ss。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 批量操作成功的数量。 **取值范围：** 不涉及。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// **参数解释：** 批量操作失败的数量。 **取值范围：** 不涉及。
	ErrorCount *int32 `json:"error_count,omitempty"`

	// 批量操作任务列表。
	ErrorItems     *[]ShowBatchOperationTaskErrorItem `json:"error_items,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowBatchOperationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchOperationTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchOperationTaskResponse", string(data)}, " ")
}
