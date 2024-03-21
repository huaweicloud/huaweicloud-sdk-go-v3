package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BusinessCatalogTreeNode 业务资产目录树
type BusinessCatalogTreeNode struct {

	// 业务资产guid
	BusinessCatalogGuid *string `json:"business_catalog_guid,omitempty"`

	// 业务资产名称
	BusinessCatalogName *string `json:"business_catalog_name,omitempty"`

	// 业务资产英文名称
	BusinessCatalogNameEng *string `json:"business_catalog_name_eng,omitempty"`

	// 业务资产级别
	Level *string `json:"level,omitempty"`

	// 业务资产级唯一限定名称
	QualifiedName *string `json:"qualified_name,omitempty"`

	// 序数
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 子级业务资产列表
	ChildNodes *[]BusinessCatalogTreeNode `json:"child_nodes,omitempty"`

	// 逻辑实体列表
	LogicEntityNodes *[]LogicEntityNodes `json:"logic_entity_nodes,omitempty"`
}

func (o BusinessCatalogTreeNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCatalogTreeNode struct{}"
	}

	return strings.Join([]string{"BusinessCatalogTreeNode", string(data)}, " ")
}
