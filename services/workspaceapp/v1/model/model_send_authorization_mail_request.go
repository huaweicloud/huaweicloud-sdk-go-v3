package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendAuthorizationMailRequest Request Object
type SendAuthorizationMailRequest struct {

	// 语言： - zh-cn：中文 - en-us：英文 - fr-fr: 法文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ResendAuthorizationMailReq `json:"body,omitempty"`
}

func (o SendAuthorizationMailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendAuthorizationMailRequest struct{}"
	}

	return strings.Join([]string{"SendAuthorizationMailRequest", string(data)}, " ")
}
