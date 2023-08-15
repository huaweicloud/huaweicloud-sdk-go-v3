package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlConfigurationRequest Request Object
type SwitchGaussMySqlConfigurationRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigurationId string `json:"configuration_id"`

	Body *ApplyConfigurationRequestBody `json:"body,omitempty"`
}

func (o SwitchGaussMySqlConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlConfigurationRequest struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlConfigurationRequest", string(data)}, " ")
}
