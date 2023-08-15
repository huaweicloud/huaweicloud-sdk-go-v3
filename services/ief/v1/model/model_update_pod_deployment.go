package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePodDeployment 应用部署全量替换修改，应用更新时调用
type UpdatePodDeployment struct {

	// 应用部署副本数
	Replicas *int32 `json:"replicas,omitempty"`

	Template *PodRequest `json:"template"`
}

func (o UpdatePodDeployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePodDeployment struct{}"
	}

	return strings.Join([]string{"UpdatePodDeployment", string(data)}, " ")
}
