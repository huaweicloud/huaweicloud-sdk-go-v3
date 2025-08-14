package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyEmailResponse Response Object
type VerifyEmailResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o VerifyEmailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyEmailResponse struct{}"
	}

	return strings.Join([]string{"VerifyEmailResponse", string(data)}, " ")
}
