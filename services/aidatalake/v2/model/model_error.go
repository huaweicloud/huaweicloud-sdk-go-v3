package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Error **参数解释**：失败时返回的文本消息内容，包含错误码、错误描述。
type Error struct {

	// **参数解释**：失败状态时的错误编码。 **取值范围**：支持英文字母、数字、小数点及下划线，长度为0~36个字符。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：失败状态时的错误信息。 **取值范围**：支持中文字符、英文字母、数字及特殊字符，长度为0~1024个字符。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o Error) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Error struct{}"
	}

	return strings.Join([]string{"Error", string(data)}, " ")
}
