package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DictionaryExtendTwo 字典扩展字段2 - 字符集：中文、英文字母、数字、下划线和空格
type DictionaryExtendTwo struct {
}

func (o DictionaryExtendTwo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryExtendTwo struct{}"
	}

	return strings.Join([]string{"DictionaryExtendTwo", string(data)}, " ")
}
