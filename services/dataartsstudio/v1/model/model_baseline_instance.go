package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaselineInstance struct {

	// 基线实例ID。
	Id *string `json:"id,omitempty"`

	// 基线明细。
	Baseline *interface{} `json:"baseline,omitempty"`

	// 基线ID。
	BaselineId *string `json:"baseline_id,omitempty"`

	// 基线任务名称。
	BaselineName *string `json:"baseline_name,omitempty"`

	// 版本号。
	BaselineVersion *int32 `json:"baseline_version,omitempty"`

	// 优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 基线任务实例作业依赖图（包含JOB_ID+JOB名称+版本号+是否为关键路径节点）。
	Dag *string `json:"dag,omitempty"`

	// 基线实例状态。
	Status *string `json:"status,omitempty"`

	// 基线实例余量，单位为s。
	Buffer *int64 `json:"buffer,omitempty"`

	// 预计完成时间戳，单位毫秒。
	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	// 实例预警时间戳，单位毫秒。
	ExpectTime *int64 `json:"expect_time,omitempty"`

	// 基线实例是否完成。
	FinishStatus *string `json:"finish_status,omitempty"`

	// 基线实例开始时间戳，单位毫秒。
	StartTime *int64 `json:"start_time,omitempty"`

	// 基线实例结束时间戳，单位毫秒，finish_status（基线实例完成状态）为FINISH（已完成）时，返回基线实例的完成时间戳。
	EndTime *int64 `json:"end_time,omitempty"`

	// 运行时间戳，单位毫秒。
	ExecuteTime *int64 `json:"execute_time,omitempty"`

	// 错误编码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 任务实例信息。
	TaskInstances *interface{} `json:"task_instances,omitempty"`

	// 责任人用户ID。
	OwnerId *string `json:"owner_id,omitempty"`

	// 责任人用户名称。
	OwnerName *string `json:"owner_name,omitempty"`

	// 责任人租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 责任人租户名称。
	DomainName *string `json:"domain_name,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// DataArts Studio实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 工作空间ID。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 首次告警时间戳，单位毫秒。
	FirstAlarmTime *int64 `json:"first_alarm_time,omitempty"`

	// 最后告警时间戳，单位毫秒。
	LastAlarmTime *int64 `json:"last_alarm_time,omitempty"`

	// 实例承诺时间戳，单位毫秒。
	SlaTime *int64 `json:"sla_time,omitempty"`

	// 处理时间戳，单位毫秒。
	ProcessTime *int64 `json:"process_time,omitempty"`

	// 恢复时间戳，单位毫秒。
	RecoverTime *int64 `json:"recover_time,omitempty"`

	// 忽略时间戳，单位毫秒。
	IgnoreTime *int64 `json:"ignore_time,omitempty"`

	// 处理时长，设置处理时间所需要的时间，设置后在该时间段内将暂停事件报警，事件的处理操作记录会被记录。
	ProcessBuffer *int64 `json:"process_buffer,omitempty"`

	// 创建时间的天数，表示一年的第几天。
	CreateDay *int32 `json:"create_day,omitempty"`

	// 实例类型。
	InstanceType *string `json:"instance_type,omitempty"`

	// 基线处理人用户ID。
	ProcessUserId *string `json:"process_user_id,omitempty"`

	// 基线处理人用户名称。
	ProcessUserName *string `json:"process_user_name,omitempty"`
}

func (o BaselineInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaselineInstance struct{}"
	}

	return strings.Join([]string{"BaselineInstance", string(data)}, " ")
}
