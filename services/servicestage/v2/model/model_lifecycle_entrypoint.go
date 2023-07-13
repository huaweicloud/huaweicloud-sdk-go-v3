package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LifecycleEntrypoint 生命周期
type LifecycleEntrypoint struct {

	// 执行命令行
	Command *[]string `json:"command,omitempty"`

	// 运行参数
	Args *[]string `json:"args,omitempty"`
}

func (o LifecycleEntrypoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleEntrypoint struct{}"
	}

	return strings.Join([]string{"LifecycleEntrypoint", string(data)}, " ")
}
