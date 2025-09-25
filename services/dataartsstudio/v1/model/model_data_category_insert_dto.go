package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataCategoryInsertDto struct {

	// 数据分类名称。
	CategoryName string `json:"category_name"`

	// 数据分类描述。
	Description *string `json:"description,omitempty"`

	// 父数据分类id，通过调用查询数据分类列表接口获取，最外层数据分类的父数据分类id为-1。默认为-1
	ParentId *string `json:"parent_id,omitempty"`
}

func (o DataCategoryInsertDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataCategoryInsertDto struct{}"
	}

	return strings.Join([]string{"DataCategoryInsertDto", string(data)}, " ")
}
