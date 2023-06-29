package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartResSceneJobsRequest Request Object
type StartResSceneJobsRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 资源id，可以为数据源id或场景id。
	ResourceId string `json:"resource_id"`

	// 动作类型： - START，启动 - STOP，停止
	Action *string `json:"action,omitempty"`
}

func (o StartResSceneJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResSceneJobsRequest struct{}"
	}

	return strings.Join([]string{"StartResSceneJobsRequest", string(data)}, " ")
}
