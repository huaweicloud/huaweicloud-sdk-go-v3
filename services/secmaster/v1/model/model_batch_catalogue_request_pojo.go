package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCatalogueRequestPojo 批量更新目录请求体
type BatchCatalogueRequestPojo struct {

	// 是否显示面包屑导航
	IsNavigation *bool `json:"is_navigation,omitempty"`

	// 是否显示名片区
	IsCardArea *bool `json:"is_card_area,omitempty"`

	// 落地页
	LandingPage *string `json:"landing_page,omitempty"`

	// 一级目录中文别名
	ParentAliasZh *string `json:"parent_alias_zh,omitempty"`

	// 一级目录中文别名
	ParentAliasEn *string `json:"parent_alias_en,omitempty"`

	// 一级目录名称
	ParentCatalogue *string `json:"parent_catalogue,omitempty"`

	// 目录详情列表
	CatalogueInfos *[]CatalogueBatchPojo `json:"catalogue_infos,omitempty"`
}

func (o BatchCatalogueRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCatalogueRequestPojo struct{}"
	}

	return strings.Join([]string{"BatchCatalogueRequestPojo", string(data)}, " ")
}
