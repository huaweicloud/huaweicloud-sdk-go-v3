package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendAuthorizedMailRequest Request Object
type SendAuthorizedMailRequest struct {

	// 语言： - zh-cn：中文 - en-us：英文 - fr-fr: 法文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ResendAuthorizedMailReq `json:"body,omitempty"`
}

func (o SendAuthorizedMailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendAuthorizedMailRequest struct{}"
	}

	return strings.Join([]string{"SendAuthorizedMailRequest", string(data)}, " ")
}
