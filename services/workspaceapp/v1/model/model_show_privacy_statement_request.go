package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivacyStatementRequest Request Object
type ShowPrivacyStatementRequest struct {

	// 语言 - zh-cn：中文 - en-us：英文 - fr-fr: 法文
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowPrivacyStatementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivacyStatementRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivacyStatementRequest", string(data)}, " ")
}
