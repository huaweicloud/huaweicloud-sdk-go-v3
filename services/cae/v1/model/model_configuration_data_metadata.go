package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationDataMetadata 组件配置元数据。
type ConfigurationDataMetadata struct {

	// 附加参数。
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o ConfigurationDataMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDataMetadata struct{}"
	}

	return strings.Join([]string{"ConfigurationDataMetadata", string(data)}, " ")
}
