package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorMessage **参数解释**：失败时返回的文本消息内容，包含错误码、错误描述和解决方案。
type ErrorMessage struct {

	// **参数解释**：错误码。 **取值范围**：长度为8~36个字符。
	ErrorCode string `json:"error_code"`

	// **参数解释**：错误描述。 **取值范围**：长度为2~512个字符。
	ErrorMsg string `json:"error_msg"`

	// **参数解释**：解决方案。 **取值范围**：长度为0~4096个字符。
	SolutionMsg *string `json:"solution_msg,omitempty"`
}

func (o ErrorMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorMessage struct{}"
	}

	return strings.Join([]string{"ErrorMessage", string(data)}, " ")
}
