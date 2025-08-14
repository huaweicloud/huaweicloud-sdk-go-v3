package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterEventResponseInfo 安全事件响应
type ClusterEventResponseInfo struct {

	// 阻断动作
	Action *string `json:"action,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群Id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 事件名称
	EventName *string `json:"event_name,omitempty"`

	// 事件唯一标识
	EventClassId *string `json:"event_class_id,omitempty"`

	// 事件id
	EventId *string `json:"event_id,omitempty"`

	// 事件类型
	EventType *int32 `json:"event_type,omitempty"`

	// 事件内容
	EventContent *string `json:"event_content,omitempty"`

	// 处理状态，包含如下:   - unhandled：未处理   - handled：已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略ID
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
