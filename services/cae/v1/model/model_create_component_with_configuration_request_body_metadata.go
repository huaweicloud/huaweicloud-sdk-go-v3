package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentWithConfigurationRequestBodyMetadata 请求数据。
type CreateComponentWithConfigurationRequestBodyMetadata struct {

	// 组件名称。
	Name string `json:"name"`

	// 创建组件请求体附加参数，当前只支持version参数，此参数必填。
	Annotations map[string]string `json:"annotations"`
}

func (o CreateComponentWithConfigurationRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentWithConfigurationRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateComponentWithConfigurationRequestBodyMetadata", string(data)}, " ")
}
