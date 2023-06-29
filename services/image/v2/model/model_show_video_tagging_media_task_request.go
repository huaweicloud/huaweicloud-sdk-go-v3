package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVideoTaggingMediaTaskRequest Request Object
type ShowVideoTaggingMediaTaskRequest struct {

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowVideoTaggingMediaTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoTaggingMediaTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoTaggingMediaTaskRequest", string(data)}, " ")
}
