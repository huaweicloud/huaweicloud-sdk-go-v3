package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerOperationError struct {

	// **参数解释**：错误编码。 **取值范围**：长度为[8,36]个字符。
	Code *string `json:"code,omitempty"`

	// **参数解释**：错误信息。 **取值范围**：长度为[2,512]个字符。
	Message *string `json:"message,omitempty"`
}

func (o ServerOperationError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerOperationError struct{}"
	}

	return strings.Join([]string{"ServerOperationError", string(data)}, " ")
}
