package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DictionaryName 字典名称 - 字符集：中文、英文字母、数字、下划线和空格 - 约束：实例下唯一
type DictionaryName struct {
}

func (o DictionaryName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryName struct{}"
	}

	return strings.Join([]string{"DictionaryName", string(data)}, " ")
}
