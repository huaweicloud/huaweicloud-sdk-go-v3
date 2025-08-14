package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerBaseInfo 容器基本信息
type ContainerBaseInfo struct {

	// **参数解释**: ID **取值范围**: 字符长度0-255位
	Id *string `json:"id,omitempty"`

	// **参数解释**: 区域 **取值范围**: 字符长度0-255位
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释**: 容器ID **取值范围**: 字符长度0-255位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 容器名称 **取值范围**: 字符长度0-255位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-255位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 容器状态 **取值范围**: - Running：运行中。 - Terminated：终止。 - Waiting：等待。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 取值0-4071095999000
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: cpu限制 **取值范围**: 字符长度0-64位
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释**: 内存限制 **取值范围**: 字符长度0-64位
	MemoryLimit *string `json:"memory_limit,omitempty"`

	// **参数解释**: 重启次数 **取值范围**: 取值0-20
	RestartCount *int32 `json:"restart_count,omitempty"`

	// **参数解释**: 所属pod名称 **取值范围**: 字符长度0-64位
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**: 所属集群 **取值范围**: 字符长度0-64位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: 集群id **取值范围**: 字符长度0-64位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 集群类型 **取值范围**: - k8s：原生集群。 - cce：CCE集群。 - ali：阿里云集群。 - tencent：腾讯云集群。 - azure：微软云集群。 - aws：亚马逊集群。 - self_built_hw：华为云自建集群。 - self_built_idc：IDC自建集群。
	ClusterType *string `json:"cluster_type,omitempty"`

	// **参数解释**: 是否有风险 **取值范围**: true和false，true代表有风险，false代表无风险
	Risky *bool `json:"risky,omitempty"`

	// **参数解释**: 低危风险数量 **取值范围**: 取值0-2147483647
	LowRisk *int32 `json:"low_risk,omitempty"`

	// **参数解释**: 中危风险数量 **取值范围**: 取值0-2147483647
	MediumRisk *int32 `json:"medium_risk,omitempty"`

	// **参数解释**: 高危风险数量 **取值范围**: 取值0-2147483647
	HighRisk *int32 `json:"high_risk,omitempty"`

	// **参数解释**: 致命风险数量 **取值范围**: 取值0-2147483647
	FatalRisk *int32 `json:"fatal_risk,omitempty"`
}

func (o ContainerBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerBaseInfo struct{}"
	}

	return strings.Join([]string{"ContainerBaseInfo", string(data)}, " ")
}
