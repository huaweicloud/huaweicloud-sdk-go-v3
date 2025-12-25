package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdhocQueryAnalysisField struct {

	// 字段名称
	Name string `json:"name"`

	// 字段类型
	LogicalType *interface{} `json:"logical_type,omitempty"`

	// 字段别名
	Alias *string `json:"alias,omitempty"`
}

func (o AdhocQueryAnalysisField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdhocQueryAnalysisField struct{}"
	}

	return strings.Join([]string{"AdhocQueryAnalysisField", string(data)}, " ")
}
