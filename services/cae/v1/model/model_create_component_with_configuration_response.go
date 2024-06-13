package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentWithConfigurationResponse Response Object
type CreateComponentWithConfigurationResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentKindObj `json:"kind,omitempty"`

	Metadata *MetadataResponse `json:"metadata,omitempty"`

	Spec *CreateComponentSpec `json:"spec,omitempty"`

	// 配置项列表。
	Configurations *[]ConfigurationItem `json:"configurations,omitempty"`

	Status         *CreateComponentWithConfigurationResponseBodyStatus `json:"status,omitempty"`
	HttpStatusCode int                                                 `json:"-"`
}

func (o CreateComponentWithConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentWithConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateComponentWithConfigurationResponse", string(data)}, " ")
}
