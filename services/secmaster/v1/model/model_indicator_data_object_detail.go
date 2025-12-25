package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorDataObjectDetail 情报详情
type IndicatorDataObjectDetail struct {
	IndicatorType *IndicatorDataObjectDetailIndicatorType `json:"indicator_type,omitempty"`

	// 值，如：ip url domain等
	Value *string `json:"value,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	Environment *IndicatorDataObjectDetailEnvironment `json:"environment,omitempty"`

	DataSource *IndicatorDataObjectDetailDataSource `json:"data_source,omitempty"`

	// 首次发生时间
	FirstReportTime *string `json:"first_report_time,omitempty"`

	// 是否删除
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 最近发生时间
	LastReportTime *string `json:"last_report_time,omitempty"`

	// 粒度（保密等级），由高到低：1（首次发现）、2（自产数据）、3（需购买）、4（外网直接查询）
	GranularMarking *int32 `json:"granular_marking,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 威胁情报ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 是否作废
	Revoked *bool `json:"revoked,omitempty"`

	// 状态， Open--打开，Closed--关闭, Revoked--作废
	Status *string `json:"status,omitempty"`

	// 威胁度， Black--黑,White--白，Gray--灰
	Verdict *string `json:"verdict,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 置信度，取值范围是80-100
	Confidence *int32 `json:"confidence,omitempty"`
}

func (o IndicatorDataObjectDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorDataObjectDetail struct{}"
	}

	return strings.Join([]string{"IndicatorDataObjectDetail", string(data)}, " ")
}
