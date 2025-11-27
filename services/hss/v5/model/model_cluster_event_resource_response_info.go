package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterEventResourceResponseInfo 事件资源信息
type ClusterEventResourceResponseInfo struct {

	// **参数解释**: 执行动作 **取值范围**: 字符长度1-256位
	EnforcementAction *string `json:"enforcement_action,omitempty"`

	// **参数解释**: 组 **取值范围**: 字符长度1-256位
	Group *string `json:"group,omitempty"`

	// **参数解释**: 信息 **取值范围**: 字符长度1-256位
	Message *string `json:"message,omitempty"`

	// **参数解释**: 名称 **取值范围**: 字符长度1-256位
	Name *string `json:"name,omitempty"`

	// **参数解释**: 命名空间 **取值范围**: 字符长度0-256位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 版本 **取值范围**: 字符长度0-128位
	Version *string `json:"version,omitempty"`

	// **参数解释**: 资源类型 **取值范围**: 字符长度1-64位
	Kind *string `json:"kind,omitempty"`

	// **参数解释**: 资源名称 **取值范围**: 字符长度1-256位
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ClusterEventResourceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterEventResourceResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterEventResourceResponseInfo", string(data)}, " ")
}
