package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 字典编码 - 字符集：英文字母、数字、下划线和空格 - 约束：实例下唯一
type DictionaryCode struct {
}

func (o DictionaryCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryCode struct{}"
	}

	return strings.Join([]string{"DictionaryCode", string(data)}, " ")
}
