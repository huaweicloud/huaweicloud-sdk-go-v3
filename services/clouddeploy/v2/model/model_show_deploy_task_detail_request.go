package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeployTaskDetailRequest struct {
	// 部署任务id

	TaskId string `json:"task_id"`
}

func (o ShowDeployTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeployTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDeployTaskDetailRequest", string(data)}, " ")
}
