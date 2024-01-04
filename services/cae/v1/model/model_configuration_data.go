package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationData 组件配置数据。
type ConfigurationData struct {
	Spec *ConfigurationRequestDataSpec `json:"spec,omitempty"`

	Metadata *ConfigurationDataMetadata `json:"metadata,omitempty"`
}

func (o ConfigurationData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationData struct{}"
	}

	return strings.Join([]string{"ConfigurationData", string(data)}, " ")
}
