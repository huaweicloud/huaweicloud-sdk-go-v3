package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源信息
type EventResourceResponseInfo struct {

	// 租户账号ID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// Region名称
	RegionName *string `json:"region_name,omitempty" xml:"region_name"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 云主机ID
	CloudId *string `json:"cloud_id,omitempty" xml:"cloud_id"`

	// 虚拟机名称
	VmName *string `json:"vm_name,omitempty" xml:"vm_name"`

	// 虚拟机UUID
	VmUuid *string `json:"vm_uuid,omitempty" xml:"vm_uuid"`

	// 容器ID
	ContainerId *string `json:"container_id,omitempty" xml:"container_id"`

	// 镜像ID
	ImageId *string `json:"image_id,omitempty" xml:"image_id"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty" xml:"image_name"`

	// 主机属性
	HostAttr *string `json:"host_attr,omitempty" xml:"host_attr"`

	// 业务服务
	Service *string `json:"service,omitempty" xml:"service"`

	// 微服务
	MicroService *string `json:"micro_service,omitempty" xml:"micro_service"`

	// 系统CPU架构
	SysArch *string `json:"sys_arch,omitempty" xml:"sys_arch"`

	// 操作系统位数
	OsBit *string `json:"os_bit,omitempty" xml:"os_bit"`

	// 操作系统类型
	OsType *string `json:"os_type,omitempty" xml:"os_type"`

	// 操作系统名称
	OsName *string `json:"os_name,omitempty" xml:"os_name"`

	// 操作系统版本
	OsVersion *string `json:"os_version,omitempty" xml:"os_version"`
}

func (o EventResourceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResourceResponseInfo struct{}"
	}

	return strings.Join([]string{"EventResourceResponseInfo", string(data)}, " ")
}
