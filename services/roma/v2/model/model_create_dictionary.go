package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDictionary struct {
	// 字典名称 - 字符集：中文、英文字母、数字、下划线和空格 - 约束：实例下唯一

	Name string `json:"name"`
	// 字典描述

	Remark *string `json:"remark,omitempty"`
	// 字典编码 - 字符集：英文字母、数字、下划线和空格 - 约束：实例下唯一

	Code string `json:"code"`
	// 字典排序，值越小顺序越靠前

	Order *int32 `json:"order,omitempty"`
	// 字典扩展字段1 - 字符集：中文、英文字母、数字、下划线和空格

	ExtendOne *string `json:"extend_one,omitempty"`
	// 字典扩展字段2 - 字符集：中文、英文字母、数字、下划线和空格

	ExtendTwo *string `json:"extend_two,omitempty"`
	// 父字典编码,为空时代表自身就是最顶级字典

	ParentCode *string `json:"parent_code,omitempty"`
}

func (o CreateDictionary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDictionary struct{}"
	}

	return strings.Join([]string{"CreateDictionary", string(data)}, " ")
}
