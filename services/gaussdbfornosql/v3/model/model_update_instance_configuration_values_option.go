package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数值对象，用户基于默认参数模板自定义的参数值。
type UpdateInstanceConfigurationValuesOption struct {

	// 参数名称。 示例：\"concurrent_reads\":\"64\"中，key值为“concurrent_reads”。
	Key string `json:"key"`

	// 参数值。 示例：\"concurrent_reads\":\"64\"，value值为“64”。
	Value string `json:"value"`
}

func (o UpdateInstanceConfigurationValuesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationValuesOption struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationValuesOption", string(data)}, " ")
}
