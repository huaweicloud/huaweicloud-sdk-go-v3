package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerConfigsResDto struct {

	// 是否使用主机网络模式
	HostNetwork *bool `json:"host_network,omitempty"`

	// 容器端口映射值
	ContainerPortList *[]ContainerPortDto `json:"container_port_list,omitempty"`
}

func (o ContainerConfigsResDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerConfigsResDto struct{}"
	}

	return strings.Join([]string{"ContainerConfigsResDto", string(data)}, " ")
}
