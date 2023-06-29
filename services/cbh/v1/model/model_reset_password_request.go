package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPasswordRequest Request Object
type ResetPasswordRequest struct {
	Body *ResetPassword `json:"body,omitempty"`
}

func (o ResetPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequest", string(data)}, " ")
}
