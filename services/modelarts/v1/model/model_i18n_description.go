package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// I18nDescription 国际化描述。
type I18nDescription struct {

	// 国际语种。
	Language *string `json:"language,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`
}

func (o I18nDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "I18nDescription struct{}"
	}

	return strings.Join([]string{"I18nDescription", string(data)}, " ")
}
