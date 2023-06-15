package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowGaussMySqlConfigurationRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigurationId string `json:"configuration_id"`
}

func (o ShowGaussMySqlConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlConfigurationRequest", string(data)}, " ")
}
