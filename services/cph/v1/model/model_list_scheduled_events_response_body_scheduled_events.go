package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListScheduledEventsResponseBodyScheduledEvents struct {

	// 计划事件唯一标识，不超过36个字节
	EventId *string `json:"event_id,omitempty"`

	// 云手机服务器的唯一标识，不超过32个字节。
	ServerId *string `json:"server_id,omitempty"`

	// 云手机服务器名称， 不超过65字符，只支持英文字母、数字、汉字、下划线和中划线。
	ServerName *string `json:"server_name,omitempty"`

	// 云手机服务器规格名称，不超过64个字节。
	ServerModelName *string `json:"server_model_name,omitempty"`

	// 服务器状态。 - 0、1、3、4：创建中 - 2：异常 - 5：正常 - 8：冻结 - 10：关机 - 11：关机中 - 12：关机失败 - 13：开机中
	ServerState *int32 `json:"server_state,omitempty"`

	// 计划事件类型，取值范围： localdisk-recovery：本地盘换盘、 system-maintenance：系统维护
	Type *string `json:"type,omitempty"`

	// 授权类型，取值范围：maintenance：授权维修、redeploy：授权重部署
	AuthorizationType *string `json:"authorization_type,omitempty"`

	// 计划事件状态， 取值范围： inquiring: 待授权、 scheduled：待执行、 executing：执行中、 completed：执行成功、 failed：执行失败、 canceled：取消
	State *string `json:"state,omitempty"`

	// 事件发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 事件开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 事件完成时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 计划执行开始时间
	NotBefore *string `json:"not_before,omitempty"`

	// 计划执行完成时间
	NotAfter *string `json:"not_after,omitempty"`

	// 计划执行开始时间deadline
	NotBeforeDeadline *string `json:"not_before_deadline,omitempty"`

	// 计划事件描述
	Description *string `json:"description,omitempty"`
}

func (o ListScheduledEventsResponseBodyScheduledEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledEventsResponseBodyScheduledEvents struct{}"
	}

	return strings.Join([]string{"ListScheduledEventsResponseBodyScheduledEvents", string(data)}, " ")
}
