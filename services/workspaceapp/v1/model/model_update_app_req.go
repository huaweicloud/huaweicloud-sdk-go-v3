package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppReq 修改应用属性(注意非空字段的限制,操作时会使用入参的字段直接覆盖已有字段的值)。
type UpdateAppReq struct {

	// 应用名称,名称需满足如下规则: 1. 名称允许可见字符或空格，不可为全空格。 2. 长度1~64个字符。
	Name *string `json:"name,omitempty"`

	// 应用版本号。
	Version *string `json:"version,omitempty"`

	// 执行路径。
	ExecutePath *string `json:"execute_path,omitempty"`

	// 应用工作目录。
	WorkPath *string `json:"work_path,omitempty"`

	// 应用描述。
	Description *string `json:"description,omitempty"`

	// 启动命令行参数。
	CommandParam *string `json:"command_param,omitempty"`

	State *AppStateEnum `json:"state,omitempty"`

	// 是否使用沙箱模式运行，取值为： - false: 表示不以沙箱模式运行 - true: 表示以沙箱模式运行
	SandboxEnable *bool `json:"sandbox_enable,omitempty"`
}

func (o UpdateAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppReq struct{}"
	}

	return strings.Join([]string{"UpdateAppReq", string(data)}, " ")
}
