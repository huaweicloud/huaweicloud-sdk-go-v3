package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComponentConfigurationRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *ComponentConfigurationKindObj `json:"kind"`

	// 配置项列表。
	Items []ConfigurationItem `json:"items"`
}

func (o CreateComponentConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateComponentConfigurationRequestBody", string(data)}, " ")
}
