package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateCombinedPublicRecordsetsTaskResponse Response Object
type BatchCreateCombinedPublicRecordsetsTaskResponse struct {

	// **参数解释：** 批量操作任务的ID。 **取值范围：** 不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateCombinedPublicRecordsetsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCombinedPublicRecordsetsTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateCombinedPublicRecordsetsTaskResponse", string(data)}, " ")
}
