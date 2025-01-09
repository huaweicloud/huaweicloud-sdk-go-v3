package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserEventRsp struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 用户名
	Username *string `json:"username,omitempty"`

	// 企业id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 事件之间的关联id
	EventTraceId *string `json:"event_trace_id,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 事件时间，UTC时间
	EventTime *string `json:"event_time,omitempty"`

	// 操作对象类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 操作对象id
	ResourceId *string `json:"resource_id,omitempty"`

	// 操作对象名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 客户端类型
	ClientType *string `json:"client_type,omitempty"`

	// 客户端ip
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端mac地址
	ClientMac *string `json:"client_mac,omitempty"`

	// 操作用户源ip
	SourceIp *string `json:"source_ip,omitempty"`

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o UserEventRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserEventRsp struct{}"
	}

	return strings.Join([]string{"UserEventRsp", string(data)}, " ")
}
