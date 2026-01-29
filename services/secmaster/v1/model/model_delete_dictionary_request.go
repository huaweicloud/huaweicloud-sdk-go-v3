package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDictionaryRequest 删除字典请求体
type DeleteDictionaryRequest struct {

	// 字典列表
	DictList *[]DictionaryDelete `json:"dict_list,omitempty"`

	// 是否删除内置字典，默认不对外开放删除内置字典工具
	IsBuiltIn *bool `json:"is_built_in,omitempty"`
}

func (o DeleteDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDictionaryRequest struct{}"
	}

	return strings.Join([]string{"DeleteDictionaryRequest", string(data)}, " ")
}
