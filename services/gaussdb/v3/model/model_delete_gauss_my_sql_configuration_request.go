package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteGaussMySqlConfigurationRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigurationId string `json:"configuration_id"`
}

func (o DeleteGaussMySqlConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlConfigurationRequest", string(data)}, " ")
}
