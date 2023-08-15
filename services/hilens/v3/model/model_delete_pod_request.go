package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePodRequest Request Object
type DeletePodRequest struct {

	// 应用部署ID
	DeploymentId string `json:"deployment_id"`

	// 是否强制删除，为true的时候为强制删除
	ForceDelete *bool `json:"force_delete,omitempty"`

	// 实例ID
	PodId string `json:"pod_id"`
}

func (o DeletePodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePodRequest struct{}"
	}

	return strings.Join([]string{"DeletePodRequest", string(data)}, " ")
}
