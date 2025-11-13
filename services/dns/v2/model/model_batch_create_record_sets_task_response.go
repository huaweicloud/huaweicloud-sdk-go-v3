package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateRecordSetsTaskResponse Response Object
type BatchCreateRecordSetsTaskResponse struct {

	// **参数解释：** 批量创建记录集任务的ID。 **取值范围：** 不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateRecordSetsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRecordSetsTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateRecordSetsTaskResponse", string(data)}, " ")
}
