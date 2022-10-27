package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationItem struct {

	// 配置类型。
	Type string `json:"type"`

	// 配置数据
	Data *interface{} `json:"data"`
}

func (o ConfigurationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationItem struct{}"
	}

	return strings.Join([]string{"ConfigurationItem", string(data)}, " ")
}
