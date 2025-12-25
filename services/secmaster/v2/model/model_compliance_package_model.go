package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompliancePackageModel struct {

	// 遵从包的uuid
	Uuid string `json:"uuid"`

	// 遵从包的名称
	Name string `json:"name"`

	// 遵从包的版本号
	Version *string `json:"version,omitempty"`

	// 遵从包的责任人
	Owner *string `json:"owner,omitempty"`

	// 遵从包的目录列表
	SpecCatalogVoList *[]BaselineCatalogModel `json:"spec_catalog_vo_list,omitempty"`

	// 对遵从包的描述
	Description string `json:"description"`

	// 遵从包的分类
	Classify *string `json:"classify,omitempty"`

	// 遵从包适用的领域
	Areas *string `json:"areas,omitempty"`

	// 遵从包适用的区域
	Region *string `json:"region,omitempty"`

	// 表示遵从包的状态 0：未启用 1: 启用
	State int32 `json:"state"`

	// 表示遵从包的类型 0：内置 1: 自定义
	Type *int32 `json:"type,omitempty"`

	// 表示该遵从包包含的检查项个数
	CheckItemsNum *int32 `json:"check_items_num,omitempty"`

	// 表示该遵从包是否包含自动检查项
	HasAutoCheckItems *bool `json:"has_auto_check_items,omitempty"`
}

func (o CompliancePackageModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompliancePackageModel struct{}"
	}

	return strings.Join([]string{"CompliancePackageModel", string(data)}, " ")
}
