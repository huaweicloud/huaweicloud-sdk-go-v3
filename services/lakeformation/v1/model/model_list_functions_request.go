package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionsRequest Request Object
type ListFunctionsRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 函数名通配符
	FunctionNamePattern *string `json:"function_name_pattern,omitempty"`

	// 查询返回满足条件的function数量
	Limit *int32 `json:"limit,omitempty"`

	// 当前查询的起始位置
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListFunctionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionsRequest", string(data)}, " ")
}
