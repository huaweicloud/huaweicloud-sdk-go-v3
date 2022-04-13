package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CreateTemporaryAccessKeyByTokenRequestBody struct {
	Auth *TokenAuth `json:"auth"`
}

func (o CreateTemporaryAccessKeyByTokenRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByTokenRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByTokenRequestBody", string(data)}, " ")
}
