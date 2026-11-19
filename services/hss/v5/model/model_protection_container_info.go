package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectionContainerInfo 容器列表
type ProtectionContainerInfo struct {

	// **参数解释**: 容器名称 **取值范围**: 字符长度0-65535位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 容器状态 **取值范围**: 字符长度0-65535位
	ContainerStatus *string `json:"container_status,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-65535位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: cpu限制值 **取值范围**: 字符长度0-65535位
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释**: 内存限制值 **取值范围**: 字符长度0-65535位
	MemoryLimit *string `json:"memory_limit,omitempty"`

	// **参数解释**: 重启次数 **取值范围**: 字符长度0-65535位
	RestartCount *string `json:"restart_count,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ProtectionContainerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionContainerInfo struct{}"
	}

	return strings.Join([]string{"ProtectionContainerInfo", string(data)}, " ")
}
