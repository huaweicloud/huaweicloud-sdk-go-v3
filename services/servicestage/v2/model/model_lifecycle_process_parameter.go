package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LifecycleProcessParameter 启动后处理或者停止前处理参数。
type LifecycleProcessParameter struct {

	// 命令参数，适用于command类型
	Command *[]string `json:"command,omitempty"`

	// 默认为POD实例的IP地址。也可以自己指定。适用于http类型。
	Host *string `json:"host,omitempty"`

	// 端口号，适用于http类型。
	Port *int32 `json:"port,omitempty"`

	// 请求url，适用于http类型。
	Path *string `json:"path,omitempty"`
}

func (o LifecycleProcessParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleProcessParameter struct{}"
	}

	return strings.Join([]string{"LifecycleProcessParameter", string(data)}, " ")
}
