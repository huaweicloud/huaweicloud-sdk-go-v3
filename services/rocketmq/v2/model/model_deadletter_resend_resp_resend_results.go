package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeadletterResendRespResendResults struct {

	// 消息ID。
	MsgId *string `json:"msg_id,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o DeadletterResendRespResendResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadletterResendRespResendResults struct{}"
	}

	return strings.Join([]string{"DeadletterResendRespResendResults", string(data)}, " ")
}
