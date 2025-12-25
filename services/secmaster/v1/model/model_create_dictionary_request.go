package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDictionaryRequest 创建字典请求体
type CreateDictionaryRequest struct {

	// 字典列表
	DictList *[]DictionaryCreate `json:"dict_list,omitempty"`

	// 是否创建内置字典
	IsBuiltIn *bool `json:"is_built_in,omitempty"`
}

func (o CreateDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDictionaryRequest struct{}"
	}

	return strings.Join([]string{"CreateDictionaryRequest", string(data)}, " ")
}
