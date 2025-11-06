package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryAlarmRulesResponse Response Object
type ListFactoryAlarmRulesResponse struct {

	// 通知规则明细。
	AlarmRulesItemsDetails *[]AlarmRulesItemsDetails `json:"alarm_rules_items_details,omitempty"`

	// 保障作业承诺时间未完成数量。
	AssuranceMissionCommitmentTimeNotCompletedCount *int32 `json:"assurance_mission_commitment_time_not_completed_count,omitempty"`

	// 保障作业失败数量。
	AssuranceMissionFailedCount *int32 `json:"assurance_mission_failed_count,omitempty"`

	// 保障作业预警时间未完成数量。
	AssuranceMissionWarningTimeNotCompletedCount *int32 `json:"assurance_mission_warning_time_not_completed_count,omitempty"`

	// 基线作业失败数量。
	BaselineErrorCount *int32 `json:"baseline_error_count,omitempty"`

	// 作业取消数量。
	CancelJobsCount *int32 `json:"cancel_jobs_count,omitempty"`

	// 作业异常数量。
	ExceptionCount *int32 `json:"exception_count,omitempty"`

	// 修改作业数量。
	ModifyCount *int32 `json:"modify_count,omitempty"`

	// 多周期未完成数量。
	MultiPeriodUnfinishedCount *int32 `json:"multi_period_unfinished_count,omitempty"`

	// 资源繁忙作业数量。
	OverloadCount *int32 `json:"overload_count,omitempty"`

	// 失败作业恢复数量。
	RecoverJobsCount *int32 `json:"recover_jobs_count,omitempty"`

	// 作业成功数量。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 作业总数量。
	Total *int32 `json:"total,omitempty"`

	// 通知规则数量。
	TotalAll *int32 `json:"total_all,omitempty"`

	// 未完成作业数量。
	UnfinishedCount *int32 `json:"unfinished_count,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ListFactoryAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryAlarmRulesResponse", string(data)}, " ")
}
