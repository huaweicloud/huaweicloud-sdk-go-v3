package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CategoryDto struct {

	// key
	Key *string `json:"key,omitempty"`

	// 中文名
	NameZh *string `json:"name_zh,omitempty"`

	// 英文名
	NameEn *string `json:"name_en,omitempty"`

	// 二级分类
	SubCategories *[]CategoryDto `json:"sub_categories,omitempty"`
}

func (o CategoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryDto struct{}"
	}

	return strings.Join([]string{"CategoryDto", string(data)}, " ")
}
