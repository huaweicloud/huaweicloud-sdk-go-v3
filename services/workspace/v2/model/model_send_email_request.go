package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendEmailRequest Request Object
type SendEmailRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`

	Body *ResendEmailReq `json:"body,omitempty"`
}

func (o SendEmailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendEmailRequest struct{}"
	}

	return strings.Join([]string{"SendEmailRequest", string(data)}, " ")
}
