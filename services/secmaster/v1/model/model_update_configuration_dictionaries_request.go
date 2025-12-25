package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigurationDictionariesRequest Request Object
type UpdateConfigurationDictionariesRequest struct {

	// 用户当前语言环境
	XLanguage string `json:"X-Language"`

	Body *CreateDictionaryRequest `json:"body,omitempty"`
}

func (o UpdateConfigurationDictionariesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationDictionariesRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationDictionariesRequest", string(data)}, " ")
}
