package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlConfigurationResponse Response Object
type ShowGaussMySqlConfigurationResponse struct {
	Configurations *ConfigurationSummary2 `json:"configurations,omitempty"`

	// 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义的参数值。
	ParameterValues map[string]string `json:"parameter_values,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ShowGaussMySqlConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlConfigurationResponse", string(data)}, " ")
}
