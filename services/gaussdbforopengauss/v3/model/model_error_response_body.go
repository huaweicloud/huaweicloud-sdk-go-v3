package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorResponseBody struct {

	// **参数解释**: 错误码。 **取值范围**: 不涉及。
	ErrorCode string `json:"error_code"`

	// **参数解释**: 错误消息。 **取值范围**: 不涉及。
	ErrorMsg string `json:"error_msg"`
}

func (o ErrorResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorResponseBody struct{}"
	}

	return strings.Join([]string{"ErrorResponseBody", string(data)}, " ")
}
