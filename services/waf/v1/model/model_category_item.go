package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CategoryItem struct {

	// 分类标识
	Category *string `json:"category,omitempty"`

	// 分类名称
	Name *string `json:"name,omitempty"`

	// 子分类列表
	Subcategories *[]SubcategoryItem `json:"subcategories,omitempty"`
}

func (o CategoryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryItem struct{}"
	}

	return strings.Join([]string{"CategoryItem", string(data)}, " ")
}
