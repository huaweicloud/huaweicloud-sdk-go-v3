package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ValidateRocketMqConsumedMessageResponse struct {

	// 消费验证结果。
	ResendResults  *[]DeadletterResendRespResendResults `json:"resend_results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ValidateRocketMqConsumedMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRocketMqConsumedMessageResponse struct{}"
	}

	return strings.Join([]string{"ValidateRocketMqConsumedMessageResponse", string(data)}, " ")
}
