package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
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
