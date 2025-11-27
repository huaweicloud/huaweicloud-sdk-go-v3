package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerlessAssetDetailResponse Response Object
type ListServerlessAssetDetailResponse struct {

	// **参数解释**: 容器ID **取值范围**: 字符长度0-255位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 容器名称 **取值范围**: 字符长度0-255位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 容器状态 **取值范围**: - Running：运行中。 - Terminated：终止。 - Waiting：等待。
	ContainerStatus *string `json:"container_status,omitempty"`

	// **参数解释**: 负载id **取值范围**: 字符长度0-255位
	WorkloadId *string `json:"workload_id,omitempty"`

	// **参数解释**: 负载名称 **取值范围**: 字符长度0-255位
	WorkloadName *string `json:"workload_name,omitempty"`

	// **参数解释**: 负载类型 **取值范围**: 字符长度0-255位
	WorkloadType *string `json:"workload_type,omitempty"`

	// **参数解释**: 集群id **取值范围**: 字符长度0-64位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 所属集群 **取值范围**: 字符长度0-64位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: 命名空间 **取值范围**: 字符长度0-64位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: pod id **取值范围**: 字符长度0-64位
	PodId *string `json:"pod_id,omitempty"`

	// **参数解释**: 实例名称 **取值范围**: 字符长度0-255位
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**: 实例ip **取值范围**: 字符长度0-255位
	PodIp *string `json:"pod_ip,omitempty"`

	// **参数解释**: 镜像id **取值范围**: 字符长度0-255位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-255位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 取值0-4071095999000
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 标签列表 **取值范围**: 字符长度0-255位
	Labels *string `json:"labels,omitempty"`

	// **参数解释**: cpu限制 **取值范围**: 字符长度0-64位
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释**: 内存限制 **取值范围**: 字符长度0-64位
	MemoryLimit *string `json:"memory_limit,omitempty"`

	// **参数解释**: 启动命令 **取值范围**: 字符长度0-65535位
	Command *string `json:"command,omitempty"`

	// **参数解释**: 启动命令参数 **取值范围**: 字符长度0-65535位
	Args *string `json:"args,omitempty"`

	// **参数解释**: 工作目录 **取值范围**: 字符长度0-65535位
	WorkingDir *string `json:"working_dir,omitempty"`

	// **参数解释**: 端口列表 **取值范围**: 取值0-65535个端口对象
	PortInfo *[]ContainerPortInfo `json:"port_info,omitempty"`

	// **参数解释**: 存储卷配置 **取值范围**: 取值0-65535个配置对象
	MountList      *[]ContainerMountInfo `json:"mount_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListServerlessAssetDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerlessAssetDetailResponse struct{}"
	}

	return strings.Join([]string{"ListServerlessAssetDetailResponse", string(data)}, " ")
}
