package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecodeAuthorizationMessageRequest Request Object
type DecodeAuthorizationMessageRequest struct {

	// 通过临时访问密钥调用接口时，需要提供“X-Security-Token”Http头，取值为临时访问密钥的security_token字段。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	Body *DecodeAuthorizationMessageReq `json:"body,omitempty"`
}

func (o DecodeAuthorizationMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecodeAuthorizationMessageRequest struct{}"
	}

	return strings.Join([]string{"DecodeAuthorizationMessageRequest", string(data)}, " ")
}
