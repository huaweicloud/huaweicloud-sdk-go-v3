package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterEventResourceResponseInfo 事件资源信息
type ClusterEventResourceResponseInfo struct {

	// 执行动作
	EnforcementAction *string `json:"enforcement_action,omitempty"`

	// Group
	Group *string `json:"group,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 资源类型
	Kind *string `json:"kind,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ClusterEventResourceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterEventResourceResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterEventResourceResponseInfo", string(data)}, " ")
}
