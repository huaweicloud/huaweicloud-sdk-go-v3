package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsOption 所添加的标签具体内容，可批量添加标签，单个实例标签上限为20个。
type TagsOption struct {

	// 标签键。最大长度36个unicode字符，不能为null或者空字符串，不能为空格，校验和使用之前会自动过滤掉前后空格。
	Key string `json:"key"`

	// 标签值。最大长度43个unicode字符，可以为空字符串，不能为空格，校验和使用之前会自动过滤掉前后空格。
	Value string `json:"value"`
}

func (o TagsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsOption struct{}"
	}

	return strings.Join([]string{"TagsOption", string(data)}, " ")
}
