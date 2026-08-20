package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionModule 扩展模块信息
type ExtensionModule struct {

	// 基础url
	BaseUrl *string `json:"base_url,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// id
	Id *int32 `json:"id,omitempty"`

	// 扩展点
	Location *string `json:"location,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	Properties *ExtensionModuleProperties `json:"properties,omitempty"`

	// 发布商
	Publisher *string `json:"publisher,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 标签。
	Tags *[]string `json:"tags,omitempty"`

	// 插件链接地址
	UrlRelative *string `json:"url_relative,omitempty"`

	// 多版本属性列表
	PropertiesList *[]interface{} `json:"properties_list,omitempty"`

	// 摘要版本号
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// 分类。
	Categories *[]string `json:"categories,omitempty"`

	// 目标。预留字段，通常为空。
	Target *string `json:"target,omitempty"`

	// 产品线。预留字段，通常为空。
	ProductLine *string `json:"product_line,omitempty"`
}

func (o ExtensionModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionModule struct{}"
	}

	return strings.Join([]string{"ExtensionModule", string(data)}, " ")
}
