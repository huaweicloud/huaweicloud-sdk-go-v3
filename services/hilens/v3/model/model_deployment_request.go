package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentRequest struct {

	// 应用部署副本数，小于100。
	Replicas int32 `json:"replicas"`

	Template *PodRequest `json:"template"`
}

func (o DeploymentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentRequest struct{}"
	}

	return strings.Join([]string{"DeploymentRequest", string(data)}, " ")
}
