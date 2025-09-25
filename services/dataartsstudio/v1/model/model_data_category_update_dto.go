package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataCategoryUpdateDto struct {

	// 分类名称。
	CategoryName string `json:"category_name"`

	// 分类描述。
	Description *string `json:"description,omitempty"`
}

func (o DataCategoryUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataCategoryUpdateDto struct{}"
	}

	return strings.Join([]string{"DataCategoryUpdateDto", string(data)}, " ")
}
