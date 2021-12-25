package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署参数配置
type CreateAppsInDeploymentV3 struct {
	// 副本数量

	Replicas int32 `json:"replicas"`

	Template *PodRequest `json:"template"`

	Annotations *Annotations `json:"annotations,omitempty"`
}

func (o CreateAppsInDeploymentV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppsInDeploymentV3 struct{}"
	}

	return strings.Join([]string{"CreateAppsInDeploymentV3", string(data)}, " ")
}
