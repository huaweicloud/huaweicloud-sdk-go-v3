package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRefreshResultRequest Request Object
type ShowRefreshResultRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowRefreshResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRefreshResultRequest struct{}"
	}

	return strings.Join([]string{"ShowRefreshResultRequest", string(data)}, " ")
}
