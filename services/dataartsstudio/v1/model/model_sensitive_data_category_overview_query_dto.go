package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SensitiveDataCategoryOverviewQueryDto 敏感数据结果分类统计概览的查询对象
type SensitiveDataCategoryOverviewQueryDto struct {

	// 当前分类节点的根节点id,根节点的
	RootId *string `json:"root_id,omitempty"`

	// 当前分类节点的父节点id
	ParentId *string `json:"parent_id,omitempty"`

	// 分类的Id
	CategoryId *string `json:"category_id,omitempty"`

	// 分类的名称
	CategoryName *string `json:"category_name,omitempty"`

	// 分类path
	CategoryPath *string `json:"category_path,omitempty"`

	// 当前分类下的敏感字段数量
	Count *int64 `json:"count,omitempty"`

	// 当前分类的子节点
	Children *[]SensitiveDataCategoryOverviewQueryDto `json:"children,omitempty"`
}

func (o SensitiveDataCategoryOverviewQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SensitiveDataCategoryOverviewQueryDto struct{}"
	}

	return strings.Join([]string{"SensitiveDataCategoryOverviewQueryDto", string(data)}, " ")
}
