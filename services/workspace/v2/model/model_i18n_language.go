package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// I18nLanguage 国际命名。
type I18nLanguage struct {

	// 语言。
	Language *string `json:"language,omitempty"`

	// 值。
	Value *string `json:"value,omitempty"`
}

func (o I18nLanguage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "I18nLanguage struct{}"
	}

	return strings.Join([]string{"I18nLanguage", string(data)}, " ")
}
