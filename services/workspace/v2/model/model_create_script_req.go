package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScriptReq 创建脚本信息。
type CreateScriptReq struct {

	// 脚本名称。
	Name *string `json:"name,omitempty"`

	// 脚本类型：POWERSHELL/BAT/SHELL。
	Type *string `json:"type,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 脚本内容。
	Content *string `json:"content,omitempty"`

	// 脚本版本。
	Version *string `json:"version,omitempty"`
}

func (o CreateScriptReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptReq struct{}"
	}

	return strings.Join([]string{"CreateScriptReq", string(data)}, " ")
}
