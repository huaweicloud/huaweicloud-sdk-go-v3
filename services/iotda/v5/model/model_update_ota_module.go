package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOtaModule 修改OTA模块对象结构体
type UpdateOtaModule struct {

	// **参数说明**：OTA模块别名。 **取值范围**：长度不超过64，只允许中文、英文字母、数字、下划线（_）、连接符（-）、英文点（.）的组合。
	AliasName *string `json:"alias_name,omitempty"`

	// **参数说明**：用于描述模块的功能等信息。 **取值范围**：长度不超过1024。
	Description *string `json:"description,omitempty"`
}

func (o UpdateOtaModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOtaModule struct{}"
	}

	return strings.Join([]string{"UpdateOtaModule", string(data)}, " ")
}
