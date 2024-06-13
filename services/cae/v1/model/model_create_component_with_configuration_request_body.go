package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComponentWithConfigurationRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *ComponentKindObj `json:"kind"`

	Metadata *CreateComponentWithConfigurationRequestBodyMetadata `json:"metadata,omitempty"`

	Spec *CreateComponentWithConfigurationRequestBodySpec `json:"spec,omitempty"`

	// 配置项列表。
	Configurations *[]ConfigurationItem `json:"configurations,omitempty"`
}

func (o CreateComponentWithConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentWithConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateComponentWithConfigurationRequestBody", string(data)}, " ")
}
