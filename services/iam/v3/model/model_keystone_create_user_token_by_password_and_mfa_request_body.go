package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateUserTokenByPasswordAndMfaRequestBody
type KeystoneCreateUserTokenByPasswordAndMfaRequestBody struct {
	Auth *MfaAuth `json:"auth"`
}

func (o KeystoneCreateUserTokenByPasswordAndMfaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordAndMfaRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordAndMfaRequestBody", string(data)}, " ")
}
