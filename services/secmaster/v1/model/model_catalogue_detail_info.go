package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogueDetailInfo 目录详情信息
type CatalogueDetailInfo struct {

	// 目录id
	Id *string `json:"id,omitempty"`

	// 一级目录名称
	ParentCatalogue *string `json:"parent_catalogue,omitempty"`

	// 二级目录名称
	SecondCatalogue *string `json:"second_catalogue,omitempty"`

	// 是否内置
	CatalogueStatus *bool `json:"catalogue_status,omitempty"`

	// 目录地址
	CatalogueAddress *string `json:"catalogue_address,omitempty"`

	// 布局ID
	LayoutId *string `json:"layout_id,omitempty"`

	// 布局名称
	LayoutName *string `json:"layout_name,omitempty"`

	// 发布者
	PublisherName *string `json:"publisher_name,omitempty"`

	// 是否展示名片区
	IsCardArea *bool `json:"is_card_area,omitempty"`

	// 是否展示目录
	IsDisplay *bool `json:"is_display,omitempty"`

	// 是否为落地页
	IsLandingPage *bool `json:"is_landing_page,omitempty"`

	// 是否展示面包屑导航
	IsNavigation *bool `json:"is_navigation,omitempty"`

	// 一级目录英文别名
	ParentAliasEn *string `json:"parent_alias_en,omitempty"`

	// 一级目录中文别名
	ParentAliasZh *string `json:"parent_alias_zh,omitempty"`

	// 二级目录英文别名
	SecondAliasEn *string `json:"second_alias_en,omitempty"`

	// 二级目录中文别名
	SecondAliasZh *string `json:"second_alias_zh,omitempty"`
}

func (o CatalogueDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogueDetailInfo struct{}"
	}

	return strings.Join([]string{"CatalogueDetailInfo", string(data)}, " ")
}
