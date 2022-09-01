package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Plugin struct {

	// 插件属性
	Attribute *string `json:"attribute,omitempty" xml:"attribute"`

	// 插件名
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o Plugin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Plugin struct{}"
	}

	return strings.Join([]string{"Plugin", string(data)}, " ")
}
