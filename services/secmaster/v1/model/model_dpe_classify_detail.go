package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DpeClassifyDetail 分类信息
type DpeClassifyDetail struct {

	// 映射id
	Id *string `json:"id,omitempty"`

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 映射id
	MappingId *string `json:"mapping_id,omitempty"`

	// 是否直接分类
	DirectClassifier *string `json:"direct_classifier,omitempty"`

	// 映射id
	DirectClassifierTypeId *string `json:"direct_classifier_type_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者id
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 修改者id
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改者名称
	ModifierName *string `json:"modifier_name,omitempty"`

	MappingInfo *DpeInfo `json:"mapping_info,omitempty"`

	// 分类集合元素
	ClassifierItemList *[]DpeClassifyItemDetail `json:"classifier_item_list,omitempty"`
}

func (o DpeClassifyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DpeClassifyDetail struct{}"
	}

	return strings.Join([]string{"DpeClassifyDetail", string(data)}, " ")
}
