package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendUserEmailVerifyCodeRequest Request Object
type SendUserEmailVerifyCodeRequest struct {
	Body *ModifyEmailAddressDto `json:"body,omitempty"`
}

func (o SendUserEmailVerifyCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendUserEmailVerifyCodeRequest struct{}"
	}

	return strings.Join([]string{"SendUserEmailVerifyCodeRequest", string(data)}, " ")
}
