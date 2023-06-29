package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchTaskResultRequest Request Object
type ShowSearchTaskResultRequest struct {

	// 分子搜索任务ID
	TaskId string `json:"task_id"`
}

func (o ShowSearchTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSearchTaskResultRequest", string(data)}, " ")
}
