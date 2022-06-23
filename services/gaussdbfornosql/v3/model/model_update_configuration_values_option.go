package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数值对象，用户基于默认参数模板自定义的参数值。为空时不修改参数值。
type UpdateConfigurationValuesOption struct {

	// Parameter name 示例：\"concurrent_reads\":\"64\"中，key值为“concurrent_reads”。 - key不为空时，value也不可为空。
	Key string `json:"key"`

	// Parameter value 示例：\"concurrent_reads\":\"64\"，value值为“64”。
	Value string `json:"value"`
}

func (o UpdateConfigurationValuesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationValuesOption struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationValuesOption", string(data)}, " ")
}
