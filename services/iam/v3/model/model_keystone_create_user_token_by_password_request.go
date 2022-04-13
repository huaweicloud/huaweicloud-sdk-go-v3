package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type KeystoneCreateUserTokenByPasswordRequest struct {
	// 如果设置该参数，返回的响应体中将不显示catalog信息。任何非空字符串都将解释为true，并使该字段生效。

	Nocatalog *string `json:"nocatalog,omitempty"`

	Body *KeystoneCreateUserTokenByPasswordRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateUserTokenByPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordRequest", string(data)}, " ")
}
