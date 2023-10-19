package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchParameter 资产查询条件
type SearchParameter struct {

	// 查询关键字
	Query *string `json:"query,omitempty"`

	Filter *DataMapFilterCriteria `json:"filter,omitempty"`

	// 条件参数列表
	Facets *[]string `json:"facets,omitempty"`

	// 分页显示每页返回结果数。默认值100
	Limit int32 `json:"limit"`

	// 偏移量，默认值0
	Offset int32 `json:"offset"`

	// 关联关系属性
	RelationshipAttributes *[]string `json:"relationship_attributes,omitempty"`

	// 排序信息
	Sort *[]Sort `json:"sort,omitempty"`

	// 所有者
	Owner *string `json:"owner,omitempty"`

	// 是否校验权限，默认值false
	QueryPrivilege *bool `json:"query_privilege,omitempty"`
}

func (o SearchParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchParameter struct{}"
	}

	return strings.Join([]string{"SearchParameter", string(data)}, " ")
}
