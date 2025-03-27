package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConsistencyResultRequest Request Object
type UpdateConsistencyResultRequest struct {

	// 任务id
	TaskId string `json:"task_id"`

	Body *SetConsistencyResultRequestBody `json:"body,omitempty"`
}

func (o UpdateConsistencyResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsistencyResultRequest struct{}"
	}

	return strings.Join([]string{"UpdateConsistencyResultRequest", string(data)}, " ")
}
