package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 个人模板参数。
type AimPersonalTemplateParam struct {

	// 动态参数类型。1：表示文本类型。
	Type int32 `json:"type"`

	// 动态参数名称。示例：${param1}。
	Name string `json:"name"`

	// 参数示例，动态参数对应的示例，不能大于100个字符。
	Example string `json:"example"`
}

func (o AimPersonalTemplateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimPersonalTemplateParam struct{}"
	}

	return strings.Join([]string{"AimPersonalTemplateParam", string(data)}, " ")
}
