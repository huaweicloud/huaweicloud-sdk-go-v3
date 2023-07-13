package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComponentConfigurationRequestBody struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“ComponentConfiguration”，该值不可修改。
	Kind string `json:"kind"`

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
