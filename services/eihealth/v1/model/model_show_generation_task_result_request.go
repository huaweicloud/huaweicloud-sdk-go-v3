package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGenerationTaskResultRequest Request Object
type ShowGenerationTaskResultRequest struct {

	// 分子生成任务ID
	TaskId string `json:"task_id"`
}

func (o ShowGenerationTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGenerationTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowGenerationTaskResultRequest", string(data)}, " ")
}
