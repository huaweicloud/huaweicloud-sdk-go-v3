package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessEventResourceResponseInfo 资源信息（当前不展示）
type ProcessEventResourceResponseInfo struct {

	// **参数解释**： 租户账号ID **取值范围**： 字符长度1-64位
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// Region ID
	RegionName *string `json:"region_name,omitempty"`

	// **参数解释**： VPC ID **取值范围**： 字符长度1-64位
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 云主机ID **取值范围**： 字符长度1-64位
	CloudId *string `json:"cloud_id,omitempty"`

	// **参数解释**： 虚拟机名称 **取值范围**： 字符长度1-64位
	VmName *string `json:"vm_name,omitempty"`

	// **参数解释**： 虚拟机UUID **取值范围**： 字符长度1-64位
	VmUuid *string `json:"vm_uuid,omitempty"`

	// **参数解释**: 容器ID **取值范围**: 字符长度1-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 镜像ID **取值范围**： 字符长度1-64位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 镜像名称，只有容器类型的告警有 **取值范围**： 字符长度1-256位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 主机属性 **取值范围**： 字符长度1-64位
	HostAttr *string `json:"host_attr,omitempty"`

	// **参数解释**： 业务服务 **取值范围**： 字符长度1-64位
	Service *string `json:"service,omitempty"`

	// **参数解释**： 微服务 **取值范围**： 字符长度1-64位
	MicroService *string `json:"micro_service,omitempty"`

	// **参数解释**： 系统CPU架构 **取值范围**： 字符长度1-64位
	SysArch *string `json:"sys_arch,omitempty"`

	// **参数解释**： 操作系统位数 **取值范围**： 字符长度1-64位
	OsBit *string `json:"os_bit,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux。 - Windows：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 操作系统名称
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： pod uid **取值范围**： 字符长度1-64位
	PodUid *string `json:"pod_uid,omitempty"`

	// **参数解释**： pod name **取值范围**： 字符长度1-64位
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**： 名称空间 **取值范围**： 字符长度1-64位
	Namespace *string `json:"namespace,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 容器状态
	ContainerStatus *string `json:"container_status,omitempty"`

	// 系统版本
	OsVersion *string `json:"os_version,omitempty"`

	// agent版本
	AgentVersion *string `json:"agent_version,omitempty"`
}

func (o ProcessEventResourceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessEventResourceResponseInfo struct{}"
	}

	return strings.Join([]string{"ProcessEventResourceResponseInfo", string(data)}, " ")
}
