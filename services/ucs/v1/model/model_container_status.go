package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerStatus 描述容器当前状态
type ContainerStatus struct {

	// 容器的唯一名称
	Name *string `json:"name,omitempty"`

	// 当前容器状态
	State *interface{} `json:"state,omitempty"`

	// 上次终止时的状态
	LastState *string `json:"lastState,omitempty"`

	// 容器是否通过就绪检查
	Ready *bool `json:"ready,omitempty"`

	// 容器重启次数
	RestartCount *int32 `json:"restartCount,omitempty"`

	// 容器运行的镜像名称
	Image *string `json:"image,omitempty"`

	// 容器运行的镜像ID
	ImageID *string `json:"imageID,omitempty"`

	// 容器是否已经成功启动并进入稳定运行阶段
	Started *string `json:"started,omitempty"`
}

func (o ContainerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerStatus struct{}"
	}

	return strings.Join([]string{"ContainerStatus", string(data)}, " ")
}
