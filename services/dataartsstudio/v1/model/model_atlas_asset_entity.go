package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AtlasAssetEntity 资产详情
type AtlasAssetEntity struct {

	// 类型名称
	TypeName string `json:"type_name"`

	// guid
	Guid *string `json:"guid,omitempty"`

	// 版本
	Version *int32 `json:"version,omitempty"`

	// 修改时间
	UpdateTime float32 `json:"update_time,omitempty"`

	// 修改人
	UpdateUser *string `json:"update_user,omitempty"`

	// 创建时间
	CreateTime float32 `json:"create_time,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 展示
	DisplayText *string `json:"display_text,omitempty"`

	// 状态 枚举值：ACTIVE、DELETED
	Status *string `json:"status,omitempty"`

	// 分类信息
	Classifications *[]AtlasClassificationInfo `json:"classifications,omitempty"`

	// 关联任务
	Meanings *[]TermAssignmentHeader `json:"meanings,omitempty"`

	// 实体map Map<String, Object>
	RelationShipAttributes *interface{} `json:"relation_ship_attributes,omitempty"`

	// 实体map Map<String, Object>
	Attributes *interface{} `json:"attributes"`
}

func (o AtlasAssetEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtlasAssetEntity struct{}"
	}

	return strings.Join([]string{"AtlasAssetEntity", string(data)}, " ")
}
