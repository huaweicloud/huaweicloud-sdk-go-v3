package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecondaryIndex 二级索引定义，元素key为数组下标。
type SecondaryIndex struct {

	// 二级索引名称，表内唯一。
	IndexName string `bson:"index_name"`

	// 排序键字段名数组，顺序组合。
	SortKeyFields []Field `bson:"sort_key_fields"`

	// 摘要字段名数组。
	AbstractFields *[]string `bson:"abstract_fields,omitempty"`
}

func (o SecondaryIndex) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecondaryIndex struct{}"
	}

	return strings.Join([]string{"SecondaryIndex", string(data)}, " ")
}
