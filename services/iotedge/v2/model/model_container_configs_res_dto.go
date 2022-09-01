package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerConfigsResDto struct {

	// 是否使用主机网络模式
	HostNetwork *bool `json:"host_network,omitempty" xml:"host_network"`

	// 容器端口映射值
	ContainerPortList *[]ContainerPortDto `json:"container_port_list,omitempty" xml:"container_port_list"`
}

func (o ContainerConfigsResDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerConfigsResDto struct{}"
	}

	return strings.Join([]string{"ContainerConfigsResDto", string(data)}, " ")
}
