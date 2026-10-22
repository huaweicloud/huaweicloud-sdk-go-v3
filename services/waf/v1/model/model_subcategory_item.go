package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubcategoryItem struct {

	// subcategory的编码
	Subcategory *string `json:"subcategory,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o SubcategoryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubcategoryItem struct{}"
	}

	return strings.Join([]string{"SubcategoryItem", string(data)}, " ")
}
