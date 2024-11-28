package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretFunctionTemplatesResponse Response Object
type ShowSecretFunctionTemplatesResponse struct {

	// 凭据轮转函数模板。
	FunctionTemplates *string `json:"function_templates,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ShowSecretFunctionTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretFunctionTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretFunctionTemplatesResponse", string(data)}, " ")
}
