package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcQuotas struct {

	// **参数解释**： 网络ACL单方向规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	AclPolicyContainRules *int32 `json:"acl_policy_contain_rules,omitempty"`

	// **参数解释**： IP地址组数配额。 **取值范围**： 整数，-1表示此配额未上线。
	AddressGroup *int32 `json:"address_group,omitempty"`

	// **参数解释**： IP地址组包含的IP地址集数配额。 **取值范围**： 整数，-1表示此配额未上线。
	AddressGroupContainIpset *int32 `json:"address_group_contain_ipset,omitempty"`

	// **参数解释**： CloudDCN场景的网络ACL数配额。 **取值范围**： 整数，-1表示此配额未上线。
	ClouddcnFirewallGroup *int32 `json:"clouddcn_firewall_group,omitempty"`

	// **参数解释**： CloudDCN场景的网络ACL单方向规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	ClouddcnAclPolicyContainRules *int32 `json:"clouddcn_acl_policy_contain_rules,omitempty"`

	// **参数解释**： 边缘网关数配额。 **取值范围**： 整数，-1表示此配额未上线。
	EdgeGateway *int32 `json:"edge_gateway,omitempty"`

	// **参数解释**： 弹性网卡关联的安全组数配额。 **取值范围**： 整数，-1表示此配额未上线。
	EniContainSecgroup *int32 `json:"eni_contain_secgroup,omitempty"`

	// **参数解释**： 网络ACL中配置了IP地址组或不连续端口的IPv4规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	FirewallPolicyContainIpv4CompositeRules *int32 `json:"firewall_policy_contain_ipv4_composite_rules,omitempty"`

	// **参数解释**： 网络ACL中配置了IP地址组或不连续端口的IPv6规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	FirewallPolicyContainIpv6CompositeRules *int32 `json:"firewall_policy_contain_ipv6_composite_rules,omitempty"`

	// **参数解释**： 云转发策略数配额。 **取值范围**： 整数，-1表示此配额未上线。
	ForwardPolicy *int32 `json:"forward_policy,omitempty"`

	// **参数解释**： 云转发策略关联端口数配额。 **取值范围**： 整数，-1表示此配额未上线。
	ForwardPolicyContainPorts *int32 `json:"forward_policy_contain_ports,omitempty"`

	// **参数解释**： 云转发规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	ForwardRule *int32 `json:"forward_rule,omitempty"`

	// **参数解释**： 对等连接数配额。 **取值范围**： 整数，-1表示此配额未上线。
	Peering *int32 `json:"peering,omitempty"`

	// **参数解释**： 一个physical_network下支持创建的roce网络数配额。 **取值范围**： 整数，-1表示此配额未上线。
	RoceClusterContainNetworks *int32 `json:"roce_cluster_contain_networks,omitempty"`

	// **参数解释**： 路由表支持系统路由数配额。 **取值范围**： 整数，-1表示此配额未上线。
	RtbAllowSysCidrRoutes *int32 `json:"rtb_allow_sys_cidr_routes,omitempty"`

	// **参数解释**： 共享带宽组数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SharedBandwidth *int32 `json:"shared_bandwidth,omitempty"`

	// **参数解释**： 单个共享带宽关联的弹性公网IP数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SharedBandwidthContainPublicip *int32 `json:"shared_bandwidth_contain_publicip,omitempty"`

	// **参数解释**： 单个安全组包含的规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SecgroupContainRules *int32 `json:"secgroup_contain_rules,omitempty"`

	// **参数解释**： 单个安全组关联的网卡数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SecgroupReferredByNic *int32 `json:"secgroup_referred_by_nic,omitempty"`

	// **参数解释**： 安全组数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SecurityGroup *int32 `json:"security_group,omitempty"`

	// **参数解释**： 安全组中配置了IP地址组、不连续端口或远端安全组的IPv4出方向规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SecurityGroupContainEgressIpv4CompositeRules *int32 `json:"security_group_contain_egress_ipv4_composite_rules,omitempty"`

	// **参数解释**： 安全组中配置了IP地址组、不连续端口或远端安全组的IPv6出方向规则数配。 **取值范围**： 整数，-1表示此配额未上线。
	SecurityGroupContainEgressIpv6CompositeRules *int32 `json:"security_group_contain_egress_ipv6_composite_rules,omitempty"`

	// **参数解释**： 安全组中配置了IP地址组、不连续端口或远端安全组的IPv4入方向规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SecurityGroupContainIngressIpv4CompositeRules *int32 `json:"security_group_contain_ingress_ipv4_composite_rules,omitempty"`

	// **参数解释**： 安全组中配置了IP地址组、不连续端口或远端安全组的IPv6入方向规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SecurityGroupContainIngressIpv6CompositeRules *int32 `json:"security_group_contain_ingress_ipv6_composite_rules,omitempty"`

	// **参数解释**： 安全组规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SecurityGroupRule *int32 `json:"security_group_rule,omitempty"`

	// **参数解释**： 单个子网包含的弹性网卡数配额。 **取值范围**： 整数，-1表示此配额未上线。
	SubnetContainEni *int32 `json:"subnet_contain_eni,omitempty"`

	// **参数解释**： 流量镜像筛选条件数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorFilter *int32 `json:"traffic_mirror_filter,omitempty"`

	// **参数解释**： 单个流量镜像筛选条件被镜像会话引用数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorFilterReferredBySession *int32 `json:"traffic_mirror_filter_referred_by_session,omitempty"`

	// **参数解释**： 流量镜像筛选条件单方向规则数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorFilterContainRulesEachDirection *int32 `json:"traffic_mirror_filter_contain_rules_each_direction,omitempty"`

	// **参数解释**： 流量镜像会话数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorSession *int32 `json:"traffic_mirror_session,omitempty"`

	// **参数解释**： 单个流量镜像会话关联的镜像源数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorSessionContainSources *int32 `json:"traffic_mirror_session_contain_sources,omitempty"`

	// **参数解释**： 单个镜像源被流量镜像会话引用次数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorSourceReferredBySession *int32 `json:"traffic_mirror_source_referred_by_session,omitempty"`

	// **参数解释**： 私网弹性负载均衡类型的镜像目的被会话引用次数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorTargetElbReferredBySession *int32 `json:"traffic_mirror_target_elb_referred_by_session,omitempty"`

	// **参数解释**： 弹性网卡类型的镜像目的被会话引用次数配额。 **取值范围**： 整数，-1表示此配额未上线。
	TrafficMirrorTargetEniReferredBySession *int32 `json:"traffic_mirror_target_eni_referred_by_session,omitempty"`

	// **参数解释**： 虚拟IP数配额。 **取值范围**： 整数，-1表示此配额未上线。
	Vip *int32 `json:"vip,omitempty"`

	// **参数解释**： 子网数配额。 **取值范围**： 整数，-1表示此配额未上线，请通过[v1查询配额接口](vpc_quota_0001.xml)。
	Virsubnet *int32 `json:"virsubnet,omitempty"`

	// **参数解释**： 单子网包含IPv4子网预留网段数配额。 **取值范围**： 整数，-1表示此配额未上线。
	VirsubnetContainIpv4CidrReservations *int32 `json:"virsubnet_contain_ipv4_cidr_reservations,omitempty"`

	// **参数解释**： 单子网包含IPv6子网预留网段数配额。 **取值范围**： 整数，-1表示此配额未上线。
	VirsubnetContainIpv6CidrReservations *int32 `json:"virsubnet_contain_ipv6_cidr_reservations,omitempty"`

	// **参数解释**： VPC数配额。 **取值范围**： 整数，-1表示此配额未上线，请通过[v1查询配额接口](vpc_quota_0001.xml)。
	Vpc *int32 `json:"vpc,omitempty"`

	// **参数解释**： 单个VPC包含弹性网卡数配额。 **取值范围**： 整数，-1表示此配额未上线。
	VpcContainEni *int32 `json:"vpc_contain_eni,omitempty"`

	// **参数解释**： 单个VPC包含虚拟IP数配额。 **取值范围**： 整数，-1表示此配额未上线。
	VpcContainVip *int32 `json:"vpc_contain_vip,omitempty"`

	// **参数解释**： 单个VPC包含子网数配额。 **取值范围**： 整数，-1表示此配额未上线。
	VpcContainVirsubnet *int32 `json:"vpc_contain_virsubnet,omitempty"`
}

func (o VpcQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcQuotas struct{}"
	}

	return strings.Join([]string{"VpcQuotas", string(data)}, " ")
}
