package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartDeploymentsPodRequest Request Object
type RestartDeploymentsPodRequest struct {

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	// 应用实例名称
	PodName string `json:"pod_name"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o RestartDeploymentsPodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartDeploymentsPodRequest struct{}"
	}

	return strings.Join([]string{"RestartDeploymentsPodRequest", string(data)}, " ")
}
