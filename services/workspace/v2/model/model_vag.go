package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Vag struct {

	// vAG信息ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 站点ID
	SiteId *string `json:"site_id,omitempty"`

	// vAG IP，与管理节点相同的IP
	VagIp *string `json:"vag_ip,omitempty"`

	// vAG内网IP，HDA往这个IP上报心跳
	PrivateIp *string `json:"private_ip,omitempty"`

	// vAG内网IPv6
	PrivateIpv6 *string `json:"private_ipv6,omitempty"`

	// vAG端口ID，与管理节点相同的端口的ID
	VagPortId *string `json:"vag_port_id,omitempty"`

	// SSH用户的名称，固定为gandalf
	SshUser *string `json:"ssh_user,omitempty"`

	// SSH用户的密码
	SshPwd *string `json:"ssh_pwd,omitempty"`

	// vAG所属ECS的ID
	VmId *string `json:"vm_id,omitempty"`

	// vAG所属机器名
	Name *string `json:"name,omitempty"`

	// vAG内部通信IP，最终租户VPC的子网IP
	InternalIp *string `json:"internal_ip,omitempty"`

	// vAG内部通信IPV6
	InternalIpv6 *string `json:"internal_ipv6,omitempty"`

	// vAG内部通信端口ID，最终租户VPC的子网中端口的ID
	InternalPortId *string `json:"internal_port_id,omitempty"`

	// 外部通信IP，可能独立的EIP
	ExternalIp *string `json:"external_ip,omitempty"`

	// 外部通信ID，可能独立的EIP ID
	ExternalId *string `json:"external_id,omitempty"`

	// root用户的名称，固定为root
	RootUser *string `json:"root_user,omitempty"`

	// root用户的密码
	RootPwd *string `json:"root_pwd,omitempty"`

	// vag操作状态
	Status *string `json:"status,omitempty"`

	// 可用分区
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建时间字符串
	CreateTimeStr *string `json:"create_time_str,omitempty"`

	// vag服务状态 NOT_USE：维护，ON_USE：启用，CANCELLATION：退服
	State *string `json:"state,omitempty"`

	// 在线用户数
	NumberOfOnlineUser *int64 `json:"number_of_online_user,omitempty"`

	// vag运行状态
	RunningStatus *string `json:"running_status,omitempty"`

	// 租户侧domainId
	DomainId *string `json:"domain_id,omitempty"`

	// vag当前版本号
	Version *string `json:"version,omitempty"`

	// vag最新版本号。
	LatestVersion *string `json:"latest_version,omitempty"`

	// wksAccessEdge版本号
	AccessEdgeVersion *string `json:"access_edge_version,omitempty"`

	// 项目是否被锁定 0是未锁定 1是锁定
	TenantLock *int32 `json:"tenant_lock,omitempty"`

	// 资源池id
	ResourcePoolId *string `json:"resource_pool_id,omitempty"`

	// agent角色，如：vag,vap4down,vap4up,authConnector
	Role *string `json:"role,omitempty"`

	// 资源池类型，public,private
	ResourcePoolType *string `json:"resource_pool_type,omitempty"`

	// 边缘sk
	EdgeSk *string `json:"edge_sk,omitempty"`

	// 是否有心跳
	HasHeartbeat *bool `json:"has_heartbeat,omitempty"`

	// VAG负载个数
	UserCount *int32 `json:"user_count,omitempty"`
}

func (o Vag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vag struct{}"
	}

	return strings.Join([]string{"Vag", string(data)}, " ")
}
