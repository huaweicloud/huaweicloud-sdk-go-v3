package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorDetail indicator detail
type CreateIndicatorDetail struct {
	DataSource *CreateAlertDataSource `json:"data_source,omitempty"`

	// 威胁度
	Verdict *string `json:"verdict,omitempty"`

	// 置信度
	Confidence *int32 `json:"confidence,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 标签
	Labels *string `json:"labels,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`

	// 粒度（保密等级），由高到低：1（首次发现）、2（自产数据）、3（需购买）、4（外网直接查询）
	GranularMarking *string `json:"granular_marking,omitempty"`

	Environment *ShowAlertRspEnvironment `json:"environment,omitempty"`

	// 是否失效
	Defanged *bool `json:"defanged,omitempty"`

	// Create time
	FirstReportTime *string `json:"first_report_time,omitempty"`

	// Update time
	LastReportTime *string `json:"last_report_time,omitempty"`

	// 指标ID
	Id *string `json:"id,omitempty"`

	IndicatorType *CreateIndicatorDetailIndicatorType `json:"indicator_type,omitempty"`

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

func (o CreateIndicatorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorDetail struct{}"
	}

	return strings.Join([]string{"CreateIndicatorDetail", string(data)}, " ")
}
