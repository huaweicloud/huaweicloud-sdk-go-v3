package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Port 功能说明：端口对象。
type Port struct {

	// **参数解释**： 端口的管理状态。 **取值范围**： true，false，默认值true。
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**： 端口所在的主机ID。 **取值范围**： 不涉及。
	BindinghostId string `json:"binding:host_id"`

	// **参数解释**： 端口的用户自定义信息。 **取值范围**： 不涉及。
	Bindingprofile *interface{} `json:"binding:profile"`

	BindingvifDetails *BindingVifDetails `json:"binding:vif_details"`

	// **参数解释**： 端口绑定的虚拟接口类型 (ovs/hw_veb等)，扩展属性。 **取值范围**： - ovs：表示使用 Open vSwitch（OVS）作为虚拟交换机 - bridge：表示使用 Linux 内核桥接（bridge）实现虚拟网络 - hw_veb：表示硬件虚拟以太网桥（Hardware Virtual Ethernet Bridge），通常用于支持 SR-IOV 的硬件网卡 - vhostuser：表示使用 vhost-user 协议（基于 Unix 域套接字）与外部虚拟交换机通信 - distributed：表示用于分布式虚拟交换机 - binding_failed：表示端口绑定失败 - unbound：表示该端口未绑定到任何网络后端
	BindingvifType string `json:"binding:vif_type"`

	// **参数解释**： 绑定的vNIC类型。 **取值范围**： - normal: 软交换。 - direct: SRIOV硬直通（不支持）。 - baremetal：用于裸金属服务器。
	BindingvnicType string `json:"binding:vnic_type"`

	// **参数解释**： 端口的创建时间。 **取值范围**： UTC时间，格式: yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// **参数解释**： 端口的最近一次更新的时间。 **取值范围**： UTC时间，格式: yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// **参数解释**： 端口的描述信息。 **取值范围**： 0-255个字符，不能包含“<”和“>”。
	Description string `json:"description"`

	// **参数解释**： 端口所属的设备ID。 **取值范围**： 带“-”的标准UUID格式。
	DeviceId string `json:"device_id"`

	// **参数解释**： 端口所属的设备名称。 **取值范围**： - network:dhcp, 表示DHCP服务 - network:router_interface_distributed, 表示子网网关地址 - compute:xxx, 表示云服务器网卡私有IP地址，其中XXX对应具体的可用区名称，例如compute:aa-bb-cc表示私有IP地址被可用区aa-bb-cc内的云服务器使用 - neutron:VIP_PORT, 表示虚拟IP地址 - neutron:LOADBALANCERV2, 表示共享型ELB - neutron:LOADBALANCERV3, 表示独享型ELB - network:endpoint_interface, 表示VPC终端节点 - network:nat_gateway, 表示NAT网关 - network:ucmp, 表示UCMP端口，为企业路由器服务所用
	DeviceOwner string `json:"device_owner"`

	// **参数解释**： 标识此端口所属云服务器的flavor。 **取值范围**： 不涉及。
	EcsFlavor string `json:"ecs_flavor"`

	// **参数解释**： 端口的资源ID。 **取值范围**： 带“-”的标准UUID格式。
	Id string `json:"id"`

	// **参数解释**： 端口所属的云服务实例ID，例如RDS实例ID。 **取值范围**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 端口所属的云服务实例类型，例如“RDS”。 **取值范围**： 不涉及。
	InstanceType string `json:"instance_type"`

	// **参数解释**： 端口的MAC地址。 **取值范围**： 不涉及
	MacAddress string `json:"mac_address"`

	// **参数解释**： 端口的名称。 **取值范围**： 默认为空，最大长度不超过255。
	Name string `json:"name"`

	// **参数解释**： 端口的安全使能标记，如果不使能，则安全组和DHCP防欺骗不生效。 **取值范围**： - true：使能端口安全。 - false：未使能端口安全。
	PortSecurityEnabled bool `json:"port_security_enabled"`

	// **参数解释**： 端口的私有IP地址。 **取值范围**： 不涉及。
	PrivateIps []PrivateIpInfo `json:"private_ips"`

	// **参数解释**： 端口所属的项目ID。 **取值范围**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 端口绑定的安全组列表。 **取值范围**： 不涉及。
	SecurityGroups []string `json:"security_groups"`

	// **参数解释**： 端口的状态。 **取值范围**： - ACTIVE：端口处于活动状态，可以正常进行网络通信。 - BUILD：端口正在创建或配置中。 - DOWN：端口处于非活动状态，不能进行网络通信。Hana 硬直通虚拟机端口状态总为DOWN。
	Status string `json:"status"`

	// **参数解释**： 端口所属的租户ID。 **取值范围**： 不涉及。
	TenantId string `json:"tenant_id"`

	// **参数解释**： 端口所在的虚拟子网ID。 **取值范围**： 带“-”的标准UUID格式。
	VirsubnetId string `json:"virsubnet_id"`

	// **参数解释**： 端口所在的VPC的ID。 **取值范围**： 带“-”的标准UUID格式。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 端口所在的VPC的租户ID。 **取值范围**： 不涉及。
	VpcTenantId string `json:"vpc_tenant_id"`

	// **参数解释**： 端口的VTEP IP地址，即虚拟隧道端点的 IP 地址。 **取值范围**： 不涉及。
	VtepIp string `json:"vtep_ip"`

	// **参数解释**： 是否使能efi，使能则表示端口支持vRoCE能力。 **取值范围**： - true：使能efi。 - false：未使能efi。
	EnableEfi bool `json:"enable_efi"`

	// **参数解释**： 端口所在子网的作用域（边缘云场景）。 **取值范围**： - center：表示作用域为中心。 - {publicBorderGroup}：表示作用域为具体的公网边界组。公网边界组限制子网的可用区范围，可关联多个边缘可用区。
	Scope string `json:"scope"`

	// **参数解释**： 端口所属的可用分区的ID。 **取值范围**： 不涉及。
	ZoneId string `json:"zone_id"`

	// **参数解释**： 端口迁移的目的节点信息，包括目的节点的binding:vif_details和binding:vif_type。 **取值范围**： 不涉及。
	BindingmigrationInfo *interface{} `json:"binding:migration_info"`

	// **参数解释**： DHCP的扩展属性。 **取值范围**： 不涉及。
	ExtraDhcpOpts []PortExtraDhcpOpt `json:"extra_dhcp_opts"`

	// **参数解释**： 边缘场景端口的位置类型。 **取值范围**： 默认值center。
	PositionType string `json:"position_type"`

	// **参数解释**： 端口绑定的实例信息。 **取值范围**： 不涉及。
	InstanceInfo *interface{} `json:"instance_info"`

	// **参数解释**： 端口的标签信息，包括标签键和标签值，可用来分类和标识资源。详情请参见Tag对象。 **取值范围**： 不涉及。
	Tags []ResponseTag `json:"tags"`

	// **参数解释**： 端口的IP/Mac对列表。 **取值范围**： - IP地址不允许为 “0.0.0.0/0”。 - 如果allowed_address_pairs配置地址池较大的IP网段（掩码小于24位），建议为该端口配置一个单独的安全组。 - 如果allowed_address_pairs的IP地址为“1.1.1.1/0”，表示关闭源目地址检查开关。 - 被绑定的云服务器网卡allowed_address_pairs的IP地址填“1.1.1.1/0”。
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs"`
}

func (o Port) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Port struct{}"
	}

	return strings.Join([]string{"Port", string(data)}, " ")
}
