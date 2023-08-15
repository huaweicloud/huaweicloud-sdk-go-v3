package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PodRequest 应用部署模板
type PodRequest struct {
	Configs *PodConfigs `json:"configs"`

	Affinity *Affinity `json:"affinity,omitempty"`

	// 应用部署init业务容器
	InitContainers *[]ContainerDef `json:"init_containers,omitempty"`

	// 应用部署业务容器
	Containers []ContainerDef `json:"containers"`
}

func (o PodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodRequest struct{}"
	}

	return strings.Join([]string{"PodRequest", string(data)}, " ")
}
