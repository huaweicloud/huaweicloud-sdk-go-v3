package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数值对象，用户基于默认参数模板自定义的参数值。默认不修改参数值。
type CreateConfigurationValuesOption struct {

	// 参数名称。 示例：\"max_connections\":\"10\"中，key值为“max_connections”。 - key为空时，不修改参数值。 - key不为空时，value也不可为空。
	Key string `json:"key"`

	// 参数值。 - 示例：\"max_connections\":\"10\"中，value值为“10”。
	Value string `json:"value"`
}

func (o CreateConfigurationValuesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationValuesOption struct{}"
	}

	return strings.Join([]string{"CreateConfigurationValuesOption", string(data)}, " ")
}
