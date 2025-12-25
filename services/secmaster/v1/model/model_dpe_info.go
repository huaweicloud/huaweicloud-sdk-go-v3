package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DpeInfo 分类信息
type DpeInfo struct {

	// 映射id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 关联分类ID
	ClassifierId *string `json:"classifier_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 完成度
	CompleteDegree *string `json:"complete_degree,omitempty"`

	// 关联实例数
	InstanceNum *string `json:"instance_num,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者id
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改者id
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改者名称
	ModifierName *string `json:"modifier_name,omitempty"`
}

func (o DpeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DpeInfo struct{}"
	}

	return strings.Join([]string{"DpeInfo", string(data)}, " ")
}
