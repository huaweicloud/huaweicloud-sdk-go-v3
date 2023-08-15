package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PodRequest struct {
	Affinity *PodAffinity `json:"affinity,omitempty"`

	// 应用部署业务容器或RPM程序包
	Apps []AppDef `json:"apps"`

	Configs *PodConfig `json:"configs"`

	// 应用部署初始化业务容器，容器部署有效。预留，暂不支持
	InitContainers *[]AppDef `json:"init_containers,omitempty"`
}

func (o PodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodRequest struct{}"
	}

	return strings.Join([]string{"PodRequest", string(data)}, " ")
}
