package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// indicator detail
type IndicatorDetail struct {

	// 指标ID
	Id *string `json:"id,omitempty"`

	// 指标名称
	Name string `json:"name"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 类型（SIMULATION,PLAYBOOK,MANUAL,INSTANCE,DATA_SOURCE）
	Type *string `json:"type,omitempty"`

	DataObject *IndicatorDataObjectDetail `json:"data_object,omitempty"`

	// workspace id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

	// 布局ID
	LayoutId *string `json:"layout_id,omitempty"`

	Dataclass *DataClassRefPojo `json:"dataclass,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o IndicatorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorDetail struct{}"
	}

	return strings.Join([]string{"IndicatorDetail", string(data)}, " ")
}
