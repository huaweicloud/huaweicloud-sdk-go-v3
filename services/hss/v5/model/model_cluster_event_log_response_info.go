package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterEventLogResponseInfo k8s集群事件日志信息
type ClusterEventLogResponseInfo struct {

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型，包含以下几种： - cce: cce集群 - ali: 阿里云集群 - tencent: 腾讯云集群 - azure: 微软云集群 - aws: 亚马逊集群 - self_built_hw: 华为云自建集群 - self_built_idc: IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`

	// 日志产生的时间
	Time *int64 `json:"time,omitempty"`

	// 触发事件的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 事件名称
	EventName *string `json:"event_name,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 事件触发的原因
	Reason *string `json:"reason,omitempty"`

	// cce集群事件的行号
	LineNum *string `json:"line_num,omitempty"`
}

func (o ClusterEventLogResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterEventLogResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterEventLogResponseInfo", string(data)}, " ")
}
