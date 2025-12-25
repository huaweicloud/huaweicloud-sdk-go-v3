package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConfigurationDictionariesRequest Request Object
type DeleteConfigurationDictionariesRequest struct {

	// 用户当前语言环境
	XLanguage string `json:"X-Language"`

	Body *DeleteDictionaryRequest `json:"body,omitempty"`
}

func (o DeleteConfigurationDictionariesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigurationDictionariesRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigurationDictionariesRequest", string(data)}, " ")
}
