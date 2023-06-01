package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SendRocketMqDlqMessageResponse struct {

	// 重发死信消息结果。
	ResendResults  *[]DeadletterResendRespResendResults `json:"resend_results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o SendRocketMqDlqMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendRocketMqDlqMessageResponse struct{}"
	}

	return strings.Join([]string{"SendRocketMqDlqMessageResponse", string(data)}, " ")
}
