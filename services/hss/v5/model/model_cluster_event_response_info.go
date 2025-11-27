package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterEventResponseInfo 安全事件响应
type ClusterEventResponseInfo struct {

	// **参数解释**: 阻断动作 **取值范围**: 字符长度1-32位
	Action *string `json:"action,omitempty"`

	// **参数解释**: 集群名称 **取值范围**: 字符长度1-64位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: 集群Id **取值范围**: 字符长度1-256位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 事件名称 **取值范围**: 字符长度1-128位
	EventName *string `json:"event_name,omitempty"`

	// **参数解释**: 事件唯一标识 **取值范围**: 字符长度1-128位
	EventClassId *string `json:"event_class_id,omitempty"`

	// **参数解释**: 事件id **取值范围**: 字符长度1-128位
	EventId *string `json:"event_id,omitempty"`

	// **参数解释**: 事件类型 **取值范围**: 最小值1000，最大值30000
	EventType *int32 `json:"event_type,omitempty"`

	// **参数解释**: 事件内容 **取值范围**: 字符长度1-128位
	EventContent *string `json:"event_content,omitempty"`

	// **参数解释**: 处理状态 **取值范围**: - unhandled：未处理。 - handled：已处理。
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 更新时间 **取值范围**: 最小值0，最大值9223372036854775807
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 项目ID **取值范围**: 字符长度1-128位
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 企业ID **取值范围**: 字符长度1-256位
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 策略名称 **取值范围**: 字符长度1-128位
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 策略ID **取值范围**: 字符长度1-128位
	PolicyId *string `json:"policy_id,omitempty"`

	ResourceInfo *ClusterEventResourceResponseInfo `json:"resource_info,omitempty"`
}

func (o ClusterEventResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterEventResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterEventResponseInfo", string(data)}, " ")
}
