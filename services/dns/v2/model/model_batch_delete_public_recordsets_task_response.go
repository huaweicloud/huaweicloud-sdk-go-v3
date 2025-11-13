package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePublicRecordsetsTaskResponse Response Object
type BatchDeletePublicRecordsetsTaskResponse struct {

	// **参数解释：** 批量操作任务的ID。 **取值范围：** 不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeletePublicRecordsetsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicRecordsetsTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicRecordsetsTaskResponse", string(data)}, " ")
}
