package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePodRequest Request Object
type DeletePodRequest struct {

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	// 应用实例名称
	PodName string `json:"pod_name"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o DeletePodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePodRequest struct{}"
	}

	return strings.Join([]string{"DeletePodRequest", string(data)}, " ")
}
