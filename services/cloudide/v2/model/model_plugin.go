package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Plugin struct {
	// 插件属性

	Attribute *string `json:"attribute,omitempty"`
	// 插件名

	Name *string `json:"name,omitempty"`
}

func (o Plugin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Plugin struct{}"
	}

	return strings.Join([]string{"Plugin", string(data)}, " ")
}
