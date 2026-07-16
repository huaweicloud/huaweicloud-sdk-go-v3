package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CombineTmsTags 多标签相同key合并value的数据结构。
type CombineTmsTags struct {

	// **参数解释**：标签的key。 **取值范围**：长度限制为128个字符，支持任意语种字母、数字、空格，以及_ . : = + - @特殊字符，但首尾不能含有空格，不能以_sys_开头。
	Key string `json:"key"`

	// **参数解释**：相同key的标签value合并后的列表。
	Values []string `json:"values"`
}

func (o CombineTmsTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CombineTmsTags struct{}"
	}

	return strings.Join([]string{"CombineTmsTags", string(data)}, " ")
}
