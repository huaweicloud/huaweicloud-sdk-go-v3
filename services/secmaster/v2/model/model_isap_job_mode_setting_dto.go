package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapJobModeSettingDto 作业模式设置数据传输对象
type IsapJobModeSettingDto struct {

	// 整型间隔时长
	BatchOvertimeInterval *int32 `json:"batch_overtime_interval,omitempty"`

	BatchOvertimeUnit *IsapTimeUnit `json:"batch_overtime_unit,omitempty"`

	// 整型间隔时长
	BatchFrequencyInterval *int32 `json:"batch_frequency_interval,omitempty"`

	BatchFrequencyUnit *IsapTimeUnit `json:"batch_frequency_unit,omitempty"`

	// 整型间隔时长
	StreamingStateTtlInterval *int32 `json:"streaming_state_ttl_interval,omitempty"`

	StreamingStateTtlUnit *IsapTimeUnit `json:"streaming_state_ttl_unit,omitempty"`

	// 整型间隔时长
	StreamingCheckpointTtlInterval *int32 `json:"streaming_checkpoint_ttl_interval,omitempty"`

	StreamingCheckpointTtlUnit *IsapTimeUnit `json:"streaming_checkpoint_ttl_unit,omitempty"`

	StreamingStartupMode *IsapJobStartupModeEnum `json:"streaming_startup_mode,omitempty"`

	// 整型间隔时长
	BatchOvertimeStrategyInterval *int32 `json:"batch_overtime_strategy_interval,omitempty"`

	BatchOvertimeStrategyUnit *IsapTimeUnit `json:"batch_overtime_strategy_unit,omitempty"`

	// 整型间隔时长
	SearchDelayInterval *int32 `json:"search_delay_interval,omitempty"`

	SearchDelayUnit *IsapTimeUnit `json:"search_delay_unit,omitempty"`

	// 整型间隔时长
	SearchFrequencyInterval *int32 `json:"search_frequency_interval,omitempty"`

	SearchFrequencyUnit *IsapTimeUnit `json:"search_frequency_unit,omitempty"`

	// 整型间隔时长
	SearchOvertimeInterval *int32 `json:"search_overtime_interval,omitempty"`

	SearchOvertimeUnit *IsapTimeUnit `json:"search_overtime_unit,omitempty"`

	// 整型间隔时长
	SearchPeriodInterval *int32 `json:"search_period_interval,omitempty"`

	SearchPeriodUnit *IsapTimeUnit `json:"search_period_unit,omitempty"`

	// UUID
	SearchTableId *string `json:"search_table_id,omitempty"`

	// 表名称
	SearchTableName *string `json:"search_table_name,omitempty"`

	FieldNotNullPolicy *IsapJobFieldNotNullPolicy `json:"field_not_null_policy,omitempty"`

	// 长整型间隔时长
	DssId *int64 `json:"dss_id,omitempty"`
}

func (o IsapJobModeSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapJobModeSettingDto struct{}"
	}

	return strings.Join([]string{"IsapJobModeSettingDto", string(data)}, " ")
}
