package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchInfo 查询条件
type SearchInfo struct {

	// 一级目录名称
	ParentCatalogue *string `json:"parent_catalogue,omitempty"`

	// 二级目录名称
	SecondCatalogue *string `json:"second_catalogue,omitempty"`

	// 是否内置
	CatalogueStatus *bool `json:"catalogue_status,omitempty"`

	// 是否内置
	CatalogueType *string `json:"catalogue_type,omitempty"`

	// 布局名称
	LayoutName *string `json:"layout_name,omitempty"`

	// 发布者
	PublisherName *string `json:"publisher_name,omitempty"`

	// 安全分析版本
	AnalysisVersion *string `json:"analysis_version,omitempty"`
}

func (o SearchInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchInfo struct{}"
	}

	return strings.Join([]string{"SearchInfo", string(data)}, " ")
}
