package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CategoryDto struct {

	// **参数解释：** 检视意见分类key。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 检视意见分类中文名。
	NameZh *string `json:"name_zh,omitempty"`

	// **参数解释：** 检视意见分类英文名。
	NameEn *string `json:"name_en,omitempty"`

	// **参数解释：** 子检视意见分类。
	SubCategories *[]CategoryDto `json:"sub_categories,omitempty"`
}

func (o CategoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryDto struct{}"
	}

	return strings.Join([]string{"CategoryDto", string(data)}, " ")
}
