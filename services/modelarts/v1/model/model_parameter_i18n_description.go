package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParameterI18nDescription 国际化描述。
type ParameterI18nDescription struct {

	// 国际语种[，可选值如下： - zh-cn（中文） - en-us（英文）](tag:hc,hk)
	Language *string `json:"language,omitempty"`

	// 国际化语种的描述信息。
	Description *string `json:"description,omitempty"`
}

func (o ParameterI18nDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterI18nDescription struct{}"
	}

	return strings.Join([]string{"ParameterI18nDescription", string(data)}, " ")
}
