package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectLogRequest Request Object
type CollectLogRequest struct {

	// 迁移任务ID
	TaskId string `json:"task_id"`

	Body *UploadLogRequestBody `json:"body,omitempty"`
}

func (o CollectLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectLogRequest struct{}"
	}

	return strings.Join([]string{"CollectLogRequest", string(data)}, " ")
}
