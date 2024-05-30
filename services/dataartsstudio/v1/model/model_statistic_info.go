package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatisticInfo 工作台统计信息，只读。
type StatisticInfo struct {
	AtomicIndex *StatisticSchema `json:"atomic_index,omitempty"`

	DerivativeIndex *StatisticSchema `json:"derivative_index,omitempty"`

	CompoundMetric *StatisticSchema `json:"compound_metric,omitempty"`

	BizIndex *StatisticSchema `json:"biz_index,omitempty"`

	Dimension *StatisticSchema `json:"dimension,omitempty"`

	ConditionGroup *StatisticSchema `json:"condition_group,omitempty"`

	TimeCondition *StatisticSchema `json:"time_condition,omitempty"`

	CommonCondition *StatisticSchema `json:"common_condition,omitempty"`

	DimensionLogicTable *StatisticSchema `json:"dimension_logic_table,omitempty"`

	FactLogicTable *StatisticSchema `json:"fact_logic_table,omitempty"`

	AggregationLogicTable *StatisticSchema `json:"aggregation_logic_table,omitempty"`

	DataStandard *StatisticSchema `json:"data_standard,omitempty"`

	TableModel *StatisticSchema `json:"table_model,omitempty"`

	LookupTable *StatisticSchema `json:"lookup_table,omitempty"`

	// 待我审核。
	PendingReview *int32 `json:"pending_review,omitempty"`

	// 我的申请。
	MyApplications *int32 `json:"my_applications,omitempty"`
}

func (o StatisticInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticInfo struct{}"
	}

	return strings.Join([]string{"StatisticInfo", string(data)}, " ")
}
