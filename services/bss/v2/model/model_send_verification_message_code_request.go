package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SendVerificationMessageCodeRequest struct {
	Body *SendVerificationCodeV2Req `json:"body,omitempty"`
}

func (o SendVerificationMessageCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVerificationMessageCodeRequest struct{}"
	}

	return strings.Join([]string{"SendVerificationMessageCodeRequest", string(data)}, " ")
}
