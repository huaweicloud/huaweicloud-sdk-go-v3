package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendVerificationMessageCodeResponse Response Object
type SendVerificationMessageCodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendVerificationMessageCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVerificationMessageCodeResponse struct{}"
	}

	return strings.Join([]string{"SendVerificationMessageCodeResponse", string(data)}, " ")
}
