package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppGroupAuthorizationRequest Request Object
type AddAppGroupAuthorizationRequest struct {

	// 语言 - zh-cn：中文 - en-us：英文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *AppGroupAuthorizeReq `json:"body,omitempty"`
}

func (o AddAppGroupAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppGroupAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"AddAppGroupAuthorizationRequest", string(data)}, " ")
}
