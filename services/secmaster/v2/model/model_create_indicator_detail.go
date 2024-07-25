package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorDetail 情报详情信息
type CreateIndicatorDetail struct {
	DataSource *CreateIndicatorDetailDataSource `json:"data_source"`

	// 威胁度
	Verdict string `json:"verdict"`

	// 置信度
	Confidence *int32 `json:"confidence,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 标签
	Labels *string `json:"labels,omitempty"`

	// 值
	Value string `json:"value"`

	// 粒度（保密等级），由高到低：1（首次发现）、2（自产数据）、3（需购买）、4（外网直接查询）
	GranularMarking string `json:"granular_marking"`

	Environment *CreateIndicatorDetailEnvironment `json:"environment"`

	// 是否失效
	Defanged bool `json:"defanged"`

	// 首次发生时间
	FirstReportTime string `json:"first_report_time"`

	// 最近发生时间
	LastReportTime *string `json:"last_report_time,omitempty"`

	// 威胁情报ID
	Id *string `json:"id,omitempty"`

	IndicatorType *CreateIndicatorDetailIndicatorType `json:"indicator_type"`

	// 威胁情报名称
	Name string `json:"name"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// workspace id
	WorkspaceId string `json:"workspace_id"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

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
