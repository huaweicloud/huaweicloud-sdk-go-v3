package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScriptReq 更新脚本信息。
type UpdateScriptReq struct {

	// 脚本名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 脚本内容。
	Content *string `json:"content,omitempty"`

	// 脚本版本。
	Version *string `json:"version,omitempty"`
}

func (o UpdateScriptReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScriptReq struct{}"
	}

	return strings.Join([]string{"UpdateScriptReq", string(data)}, " ")
}
