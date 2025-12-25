package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DictionaryDelete struct {

	// 字典id
	DictId string `json:"dict_id"`

	// 字典key
	DictKey string `json:"dict_key"`

	// 用户当前语言环境
	Language string `json:"language"`
}

func (o DictionaryDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DictionaryDelete struct{}"
	}

	return strings.Join([]string{"DictionaryDelete", string(data)}, " ")
}
