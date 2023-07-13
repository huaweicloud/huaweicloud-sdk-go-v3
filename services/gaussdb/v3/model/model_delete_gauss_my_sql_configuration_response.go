package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlConfigurationResponse Response Object
type DeleteGaussMySqlConfigurationResponse struct {

	// 参数模板ID。
	ConfigurationId *string `json:"configuration_id,omitempty"`

	// 参数模板名称。
	ConfigurationName *string `json:"configuration_name,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o DeleteGaussMySqlConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlConfigurationResponse", string(data)}, " ")
}
