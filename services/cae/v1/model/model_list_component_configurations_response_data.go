package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentConfigurationsResponseData 组件配置数据。
type ListComponentConfigurationsResponseData struct {
	Spec *ConfigurationResponseDataSpec `json:"spec,omitempty"`

	Metadata *ConfigurationDataMetadata `json:"metadata,omitempty"`
}

func (o ListComponentConfigurationsResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentConfigurationsResponseData struct{}"
	}

	return strings.Join([]string{"ListComponentConfigurationsResponseData", string(data)}, " ")
}
