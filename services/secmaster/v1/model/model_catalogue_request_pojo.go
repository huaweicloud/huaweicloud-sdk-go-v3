package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogueRequestPojo 新增、编辑目录pojo
type CatalogueRequestPojo struct {

	// 一级目录名称
	ParentCatalogue string `json:"parent_catalogue"`

	// 一级目录英文别名
	ParentAliasEn string `json:"parent_alias_en"`

	// 一级目录中文别名
	ParentAliasZh string `json:"parent_alias_zh"`

	// 二级目录名称
	SecondCatalogue string `json:"second_catalogue"`

	// 二级目录英文别名
	SecondAliasEn string `json:"second_alias_en"`

	// 二级目录中文别名
	SecondAliasZh string `json:"second_alias_zh"`

	// 目录编码
	SecondCatalogueCode string `json:"second_catalogue_code"`

	// 布局ID.
	LayoutId string `json:"layout_id"`

	// 布局名称
	LayoutName *string `json:"layout_name,omitempty"`

	// 目录地址
	CatalogueAddress string `json:"catalogue_address"`

	// 发布者
	PublisherName *string `json:"publisher_name,omitempty"`
}

func (o CatalogueRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogueRequestPojo struct{}"
	}

	return strings.Join([]string{"CatalogueRequestPojo", string(data)}, " ")
}
