package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomHooks 自定义启动脚本钩子配置。
type CustomHooks struct {
	ContainerHooks *ContainerHooks `json:"container_hooks,omitempty"`
}

func (o CustomHooks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomHooks struct{}"
	}

	return strings.Join([]string{"CustomHooks", string(data)}, " ")
}
