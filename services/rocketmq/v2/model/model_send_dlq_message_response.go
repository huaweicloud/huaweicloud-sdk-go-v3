package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SendDlqMessageResponse struct {

	// 重发死信消息结果。
	ResendResults  *[]DeadletterResendRespResendResults `json:"resend_results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o SendDlqMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDlqMessageResponse struct{}"
	}

	return strings.Join([]string{"SendDlqMessageResponse", string(data)}, " ")
}
