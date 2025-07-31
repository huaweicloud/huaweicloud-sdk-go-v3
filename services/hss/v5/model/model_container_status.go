package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerStatus 容器状态
type ContainerStatus struct {
}

func (o ContainerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerStatus struct{}"
	}

	return strings.Join([]string{"ContainerStatus", string(data)}, " ")
}
