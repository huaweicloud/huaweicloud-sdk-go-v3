package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmRulesItemsDetails struct {

	// 作业名称。
	Name *string `json:"name,omitempty"`

	// ID，与业务无关。
	Id *int64 `json:"id,omitempty"`

	// 作业ID/节点任务ID。
	NodeId *int64 `json:"node_id,omitempty"`

	// 通知类型，0:完成;1:运行异常;3:未完成; 4:资源繁忙; 5-11基线相关的告警。
	RemindType *int32 `json:"remind_type,omitempty"`

	// 租户创建的smn服务的topic名称。
	TopicName *string `json:"topic_name,omitempty"`

	// topic对应URN。
	Urn *string `json:"urn,omitempty"`

	// 节点类型，1:作业; 2: 节点任务; 3: 基线任务。
	NotifyType *int32 `json:"notify_type,omitempty"`

	// 责任人电话。
	DisplayNumber *string `json:"display_number,omitempty"`

	// 被叫方的电话。
	CalleeNumber *string `json:"callee_number,omitempty"`

	// 完成时间。
	CompleteTime *string `json:"complete_time,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 开关。
	UseFlag *int32 `json:"use_flag,omitempty"`

	// 创建人。
	CreateUser *string `json:"create_user,omitempty"`

	// 规则名称。
	RuleName *string `json:"rule_name,omitempty"`

	// 任务告警信息。
	AlarmPeriods *[]AlarmPeriod `json:"alarm_periods,omitempty"`

	// 作业目录信息。
	JobDirectory *[]DirectoryDto `json:"job_directory,omitempty"`

	// 作业ID/节点任务ID。
	NodeIdList *[]int64 `json:"node_id_list,omitempty"`

	// 作业名称/作业名称.节点任务名称。
	NodeNameList *[]string `json:"node_name_list,omitempty"`

	// 添加类型。
	AddMode *string `json:"add_mode,omitempty"`

	// 告警对象类型。
	SubjectType *string `json:"subject_type,omitempty"`

	// 告警方式。
	NotifyMethod *string `json:"notify_method,omitempty"`

	// 终端协议。
	Protocol *string `json:"protocol,omitempty"`

	// 抄送人。
	OtherPersons *string `json:"other_persons,omitempty"`

	// 最大通知次数。
	MaxSendTimes *int32 `json:"max_send_times,omitempty"`

	// 最小通知间隔，单位分钟。
	SendInterval *int32 `json:"send_interval,omitempty"`

	// 值班表id。
	DutyScheduleId *int64 `json:"duty_schedule_id,omitempty"`

	// smn配置id。
	SmnConfigId *string `json:"smn_config_id,omitempty"`
}

func (o AlarmRulesItemsDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRulesItemsDetails struct{}"
	}

	return strings.Join([]string{"AlarmRulesItemsDetails", string(data)}, " ")
}
