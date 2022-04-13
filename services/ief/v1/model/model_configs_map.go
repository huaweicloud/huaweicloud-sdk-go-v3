package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 环境变量引用配置项时使用。
type ConfigsMap struct {
	// 配置项的名称

	Name string `json:"name"`
	// 配置项的属性名

	Key string `json:"key"`
}

func (o ConfigsMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigsMap struct{}"
	}

	return strings.Join([]string{"ConfigsMap", string(data)}, " ")
}
