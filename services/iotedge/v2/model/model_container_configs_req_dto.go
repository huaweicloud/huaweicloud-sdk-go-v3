package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerConfigsReqDto struct {

	// 容器端口映射值
	ContainerPortList *[]ContainerPortDto `json:"container_port_list,omitempty"`
}

func (o ContainerConfigsReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerConfigsReqDto struct{}"
	}

	return strings.Join([]string{"ContainerConfigsReqDto", string(data)}, " ")
}
