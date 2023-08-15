package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IgnoreAdvanced 若想忽略来源于某攻击事件下指定字段的攻击，可在“高级设置”里选择指定字段进行配置，配置完成后，WAF将不再拦截指定字段的攻击事件。当时，当不检测模块为所有检测模块时，不支持配置该高级配置。
type IgnoreAdvanced struct {

	// 字段类型，支持的字段类型有：Params、Cookie、Header、Body、Multipart。   - 当选择“Params”、“Cookie”或者“Header”字段时，可以配置“全部”或根据需求配置子字段   - 当选择“Body”或“Multipart”字段时，可以配置“全部”
	Index *string `json:"index,omitempty"`

	// 指定字段类型的子字段，默认值为“全部”
	Contents *[]string `json:"contents,omitempty"`
}

func (o IgnoreAdvanced) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IgnoreAdvanced struct{}"
	}

	return strings.Join([]string{"IgnoreAdvanced", string(data)}, " ")
}
