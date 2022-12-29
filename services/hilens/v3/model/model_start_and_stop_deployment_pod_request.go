package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartAndStopDeploymentPodRequest struct {

	// 部署ID，可以在部署详情中获取指定ID
	DeploymentId string `json:"deployment_id"`

	// 节点ID, 可以在部署详情中获取指定ID
	PodId string `json:"pod_id"`

	// 操作请求，分别为，pause停止，resume启动
	Action string `json:"action"`
}

func (o StartAndStopDeploymentPodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAndStopDeploymentPodRequest struct{}"
	}

	return strings.Join([]string{"StartAndStopDeploymentPodRequest", string(data)}, " ")
}
