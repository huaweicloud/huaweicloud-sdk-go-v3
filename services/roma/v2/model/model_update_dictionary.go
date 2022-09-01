package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDictionary struct {

	// 字典名称 - 字符集：中文、英文字母、数字、下划线和空格 - 约束：实例下唯一
	Name *string `json:"name,omitempty" xml:"name"`

	// 字典描述
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 字典排序，值越小顺序越靠前
	Order *int32 `json:"order,omitempty" xml:"order"`

	// 字典扩展字段1 - 字符集：中文、英文字母、数字、下划线和空格
	ExtendOne *string `json:"extend_one,omitempty" xml:"extend_one"`

	// 字典扩展字段2 - 字符集：中文、英文字母、数字、下划线和空格
	ExtendTwo *string `json:"extend_two,omitempty" xml:"extend_two"`
}

func (o UpdateDictionary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDictionary struct{}"
	}

	return strings.Join([]string{"UpdateDictionary", string(data)}, " ")
}
