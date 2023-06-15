package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateGaussMySqlConfigurationRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigurationId string `json:"configuration_id"`

	Body *UpdateConfigurationParameterRequestBody `json:"body,omitempty"`
}

func (o UpdateGaussMySqlConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlConfigurationRequest", string(data)}, " ")
}
