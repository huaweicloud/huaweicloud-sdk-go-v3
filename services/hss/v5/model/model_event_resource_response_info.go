package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventResourceResponseInfo 资源信息
type EventResourceResponseInfo struct {

	// **参数解释**： 租户账号ID **取值范围**： 字符长度1-256位
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 项目ID **取值范围**： 字符长度1-256位
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 企业项目ID **取值范围**： 字符长度1-256位
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： Region名称 **取值范围**： 字符长度1-256位
	RegionName *string `json:"region_name,omitempty"`

	// **参数解释**： VPC ID **取值范围**： 字符长度1-256位
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 云主机ID **取值范围**： 字符长度1-256位
	CloudId *string `json:"cloud_id,omitempty"`

	// **参数解释**： 虚拟机名称 **取值范围**： 字符长度1-256位
	VmName *string `json:"vm_name,omitempty"`

	// **参数解释**： 虚拟机UUID，即主机ID **取值范围**： 字符长度1-256位
	VmUuid *string `json:"vm_uuid,omitempty"`

	// **参数解释**： 容器ID **取值范围**： 字符长度1-256位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器状态 **取值范围**： 字符长度1-256位
	ContainerStatus *string `json:"container_status,omitempty"`

	// **参数解释**： pod uid **取值范围**： 字符长度1-256位
	PodUid *string `json:"pod_uid,omitempty"`

	// **参数解释**： pod name **取值范围**： 字符长度1-256位
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**： namespace **取值范围**： 字符长度1-256位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 集群ID **取值范围**： 字符长度1-256位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 字符长度1-256位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 镜像ID **取值范围**： 字符长度1-256位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 字符长度1-256位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 主机属性 **取值范围**： 字符长度1-256位
	HostAttr *string `json:"host_attr,omitempty"`

	// **参数解释**： 业务服务 **取值范围**： 字符长度1-256位
	Service *string `json:"service,omitempty"`

	// **参数解释**： 微服务 **取值范围**： 字符长度1-256位
	MicroService *string `json:"micro_service,omitempty"`

	// **参数解释**： 系统CPU架构 **取值范围**： 字符长度1-256位
	SysArch *string `json:"sys_arch,omitempty"`

	// **参数解释**： 操作系统位数 **取值范围**： 字符长度1-256位
	OsBit *string `json:"os_bit,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： 字符长度1-256位
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 操作系统名称 **取值范围**： 字符长度1-256位
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**： 操作系统版本 **取值范围**： 字符长度1-256位
	OsVersion *string `json:"os_version,omitempty"`
}

func (o EventResourceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResourceResponseInfo struct{}"
	}

	return strings.Join([]string{"EventResourceResponseInfo", string(data)}, " ")
}
