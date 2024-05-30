package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 资源标签。
type Tag struct {

	// 键。  - 最大长度127个unicode字符。          - 不能为空，可以包含任意语种字母、数字、空格和_.：=+-@，但首尾不能含有空所有服务均可在标签输入框下拉选择同一标格，不能以_sys_开头。
	Key string `json:"key"`

	// 值列表。  - 每个值最大长度255个unicode字符。 - 不能为空，可以包含任意语种字母、、数字、空格和_.：=+-@，但首尾不能含有空格。
	Value string `json:"value"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
