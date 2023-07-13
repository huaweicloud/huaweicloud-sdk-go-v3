package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAndStopDeploymentRequest Request Object
type StartAndStopDeploymentRequest struct {

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	// 操作请求，分别为，pause停止，resume启动
	Action string `json:"action"`
}

func (o StartAndStopDeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAndStopDeploymentRequest struct{}"
	}

	return strings.Join([]string{"StartAndStopDeploymentRequest", string(data)}, " ")
}
