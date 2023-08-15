package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CategoryDetailVo struct {

	// 目录ID，根目录的ID为0
	Id *int64 `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 类型 built_in:系统内置 user-defined: 用户自定义
	Type *string `json:"type,omitempty"`

	// 父目录ID
	Pid *string `json:"pid,omitempty"`

	// 子目录
	SubCategories *[]SubCategoryDetailVo `json:"sub_categories,omitempty"`
}

func (o CategoryDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryDetailVo struct{}"
	}

	return strings.Join([]string{"CategoryDetailVo", string(data)}, " ")
}
