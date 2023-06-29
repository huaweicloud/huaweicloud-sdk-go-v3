package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSynthesisTaskResultRequest Request Object
type ShowSynthesisTaskResultRequest struct {

	// 分子合成路径规划任务ID
	TaskId string `json:"task_id"`
}

func (o ShowSynthesisTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSynthesisTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSynthesisTaskResultRequest", string(data)}, " ")
}
