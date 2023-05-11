package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户开启的自定义属性
type CustomProp struct {

	// 自定义属性的ID（API侧）
	Id string `json:"id"`

	PropDefinition *PropDefinition `json:"prop_definition,omitempty"`
}

func (o CustomProp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomProp struct{}"
	}

	return strings.Join([]string{"CustomProp", string(data)}, " ")
}
