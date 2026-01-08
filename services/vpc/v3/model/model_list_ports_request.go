package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsRequest Request Object
type ListPortsRequest struct {

	// **参数解释**： 端口的资源ID。 **取值范围**： 带“-”的标准UUID格式。
	Id *[]string `json:"id,omitempty"`

	// **参数解释**： 端口的名称。 **取值范围**： 默认为空，最大长度不超过255。
	Name *[]string `json:"name,omitempty"`

	// **参数解释**： 端口的管理状态。 **取值范围**： true，false，默认值true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**： 端口的状态。 **取值范围**： - ACTIVE：端口处- 于活动状态，可以正常进行网络通信。 - BUILD：端口正在创建或配置中。 - DOWN：端口处于非活动状态，不能进行网络通信。Hana 硬直通虚拟机端口状态总为DOWN。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 端口所在的虚拟子网ID。 **取值范围**： 带“-”的标准UUID格式。
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	// **参数解释**： 端口所属的设备ID。 **取值范围**： 带“-”的标准UUID格式。
	DeviceId *[]string `json:"device_id,omitempty"`

	// **参数解释**： 端口的MAC地址。 **取值范围**： 不涉及
	MacAddress *[]string `json:"mac_address,omitempty"`

	// **参数解释**： 端口所属的设备名称。 **取值范围**： - network:dhcp, 表示DHCP服务 - network:router_interface_distributed, 表示子网网关地址 - compute:xxx, 表示云服务器网卡私有IP地址，其中XXX对应具体的可用区名称，例如compute:aa-bb-cc表示私有IP地址被可用区aa-bb-cc内的云服务器使用 - neutron:VIP_PORT, 表示虚拟IP地址 - neutron:LOADBALANCERV2, 表示共享型ELB - neutron:LOADBALANCERV3, 表示独享型ELB - network:endpoint_interface, 表示VPC终端节点 - network:nat_gateway, 表示NAT网关 - network:ucmp, 表示UCMP端口，为企业路由器服务所用
	DeviceOwner *[]string `json:"device_owner,omitempty"`

	// **参数解释**： 端口所属的设备名称前缀。 **取值范围**： - network：过滤出device_owner前缀是network的端口，如DHCP端口。 - compute：过滤出device_owner前缀是compute的端口，如云服务器网卡。 - neutron：过滤出device_owner前缀是compute的端口，如虚拟IP地址。
	DeviceOwnerPrefixlike *string `json:"device_owner_prefixlike,omitempty"`

	// **参数解释**： 端口的描述信息。 **取值范围**： 0-255个字符，不能包含“<”和“>”。
	Description *[]string `json:"description,omitempty"`

	// **参数解释**： 端口所在的主机ID。 **取值范围**： 不涉及。
	BindinghostId *[]string `json:"binding:host_id,omitempty"`

	// **参数解释**： 端口的私有IP地址。 **取值范围**： - private_ips=ip_address={ip_address}，其中{ip_address}填IP地址，如192.168.21.22。 - private_ips=subnet_cidr_id={subnet_id}，其中{subnet_id}填IPv4子网或IPv6子网的ID，如011fc878-5521-4654-a1ad-f5b0b5820302。
	PrivateIps *[]string `json:"private_ips,omitempty"`

	// **参数解释**： 端口绑定的安全组列表。 **取值范围**： 不涉及。
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// **参数解释**： 端口所在的VPC的ID。 **取值范围**： 带“-”的标准UUID格式。
	VpcId *[]string `json:"vpc_id,omitempty"`

	// **参数解释**： 端口的IP/Mac对列表。 **取值范围**： - allowed_address_pairs=ip_address={ip_address}，其中{ip_address}填IP地址，如192.168.21.22。 - allowed_address_pairs=mac_address={mac_address}，其中{mac_address}填MAC地址，如fa:16:3e:b1:da:62。
	AllowedAddressPairs *[]string `json:"allowed_address_pairs,omitempty"`

	// **参数解释**： 端口所属的云服务实例ID，例如RDS实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 端口所属的云服务实例类型，例如“RDS”。 **取值范围**： 不涉及。
	InstanceType *string `json:"instance_type,omitempty"`
}

func (o ListPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsRequest struct{}"
	}

	return strings.Join([]string{"ListPortsRequest", string(data)}, " ")
}
