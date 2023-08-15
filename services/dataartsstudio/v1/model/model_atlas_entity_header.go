package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AtlasEntityHeader 资产详情
type AtlasEntityHeader struct {

	// 类型名称
	TypeName *string `json:"type_name,omitempty"`

	// guid
	Guid *string `json:"guid,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 展示
	DisplayText *string `json:"display_text,omitempty"`

	// 状态 枚举值：ACTIVE、DELETED
	Status *string `json:"status,omitempty"`

	ClassificationNames *[]string `json:"classification_names,omitempty"`

	// 分类信息
	Classifications *[]AtlasClassificationInfo `json:"classifications,omitempty"`

	MeaningNames *[]string `json:"meaning_names,omitempty"`

	// 关联任务
	Meanings *[]TermAssignmentHeader `json:"meanings,omitempty"`

	// 实体map Map<String, AtlasEntityHeader>
	Children *interface{} `json:"children,omitempty"`

	// 实体map Map<String, Object>
	Attributes *interface{} `json:"attributes,omitempty"`
}

func (o AtlasEntityHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtlasEntityHeader struct{}"
	}

	return strings.Join([]string{"AtlasEntityHeader", string(data)}, " ")
}
