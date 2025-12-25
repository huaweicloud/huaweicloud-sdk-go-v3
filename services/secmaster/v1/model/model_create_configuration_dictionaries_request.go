package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationDictionariesRequest Request Object
type CreateConfigurationDictionariesRequest struct {

	// 用户当前语言环境
	XLanguage string `json:"X-Language"`

	Body *CreateDictionaryRequest `json:"body,omitempty"`
}

func (o CreateConfigurationDictionariesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationDictionariesRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationDictionariesRequest", string(data)}, " ")
}
