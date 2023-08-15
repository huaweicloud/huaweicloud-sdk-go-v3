package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartResJobRequest Request Object
type StartResJobRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 资源id，可以为数据源id或场景id。
	ResourceId string `json:"resource_id"`

	// 作业id。
	JobId string `json:"job_id"`

	// 动作类型： - START，启动 - STOP，停止
	Action *string `json:"action,omitempty"`
}

func (o StartResJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResJobRequest struct{}"
	}

	return strings.Join([]string{"StartResJobRequest", string(data)}, " ")
}
