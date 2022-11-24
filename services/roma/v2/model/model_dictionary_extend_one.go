package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 字典扩展字段1 - 字符集：中文、英文字母、数字、下划线和空格
type DictionaryExtendOne struct {
}

func (o DictionaryExtendOne) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryExtendOne struct{}"
	}

	return strings.Join([]string{"DictionaryExtendOne", string(data)}, " ")
}
