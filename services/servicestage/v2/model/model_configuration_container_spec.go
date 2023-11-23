package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationContainerSpec struct {

	// 容器信息
	Containers *[]ComponentContainerParameter `json:"containers,omitempty"`

	// 工作负载类型。
	Type *string `json:"type,omitempty"`
}

func (o ConfigurationContainerSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationContainerSpec struct{}"
	}

	return strings.Join([]string{"ConfigurationContainerSpec", string(data)}, " ")
}
