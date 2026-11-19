package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubNetworkInterfacesRequest Request Object
type ListSubNetworkInterfacesRequest struct {

	// 功能说明：每页返回的个数 取值范围：0-2000
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// 功能说明：辅助弹性网卡ID，支持多ID过滤 使用场景：查询需要的多个辅助弹性网卡信息
	Id *[]string `json:"id,omitempty"`

	// 功能说明：辅助弹性网卡所属虚拟子网的ID，支持多个ID过滤 使用场景：过滤需要的单个或者多个虚拟子网下的辅助弹性网卡
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	// 功能说明：辅助弹性网卡的私有IPv4地址，支持多个地址同时过滤 使用场景：通过单个或者多个ip地址过滤查询辅助弹性网卡
	PrivateIpAddress *[]string `json:"private_ip_address,omitempty"`

	// 功能说明：辅助弹性网卡的mac地址，支持多个同时过滤 使用场景：使用mac地址精确过滤辅助弹性网卡
	MacAddress *[]string `json:"mac_address,omitempty"`

	// 功能说明：辅助弹性网卡所属的VPC_ID，支持多ID过滤 使用场景：过滤单个或多个VPC下的辅助弹性网卡信息
	VpcId *[]string `json:"vpc_id,omitempty"`

	// 功能说明：辅助弹性网卡的描述信息，支持多个同时过滤 使用场景：通过描述信息过滤辅助弹性网卡
	Description *[]string `json:"description,omitempty"`

	// 功能说明：辅助弹性网卡的宿主网卡的ID，支持多ID过滤 使用场景：过滤单个或多个宿主网卡下存在的辅助弹性网卡
	ParentId *[]string `json:"parent_id,omitempty"`

	// **参数解释**： 辅助弹性网卡是否关联EIP，可以通过此字段过滤辅助弹性网卡，不支持传入多个值进行过滤。 **取值范围**：   - true：辅助弹性网卡已关联EIP   - false：辅助弹性网卡未关联EIP **约束限制**： 不涉及。 **默认取值**： 不涉及。
	EnableEip *bool `json:"enable_eip,omitempty"`

	// **参数解释**： 辅助弹性网卡所在站点的公网出口信息，可以通过此字段过滤辅助弹性网卡。 **取值范围**： 不涉及。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	Scope *string `json:"scope,omitempty"`

	// **参数解释**： 辅助弹性网卡关联的安全组ID，可以通过此字段过滤辅助弹性网卡，支持传入多个安全组ID过滤。[获取方式请参见[查询安全组列表](vpc_apiv3_0011.xml)。](tag:hws,hws_hk,hws_test) **取值范围**： 不涉及。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// **参数解释**： 辅助弹性网卡所属企业项目的ID。可以使用该字段过滤某个企业项目下的辅助弹性网卡。 **取值范围**： - 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 - 若需要查询当前用户所有企业项目绑定的辅助弹性网卡，请传参all_granted_eps。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 辅助弹性网卡的宿主网卡所属的设备ID，可以通过此字段过滤辅助弹性网卡，支持传入多个ID过滤。 **取值范围**： 不涉及。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	ParentDeviceId *[]string `json:"parent_device_id,omitempty"`

	// **参数解释**： 辅助弹性网卡的VLAN ID，可以通过此字段过滤辅助弹性网卡，支持传入多个VLAN ID过滤。 **取值范围**： 1-4094。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	VlanId *[]int32 `json:"vlan_id,omitempty"`

	// **参数解释**： 辅助弹性网卡的状态，可以通过此字段过滤辅助弹性网卡，支持传入多个状态过滤。 **取值范围**： NORMAL：过滤出已挂载在弹性网卡上的辅助弹性网卡 UNBOUND：过滤出未挂载在弹性网卡上的辅助弹性网卡 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	State *[]string `json:"state,omitempty"`

	// **参数解释**： 辅助弹性网卡的IPv6地址，可以通过此字段过滤辅助弹性网卡，支持传入多个IP地址过滤。 **取值范围**： 不涉及。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	Ipv6IpAddress *[]string `json:"ipv6_ip_address,omitempty"`

	// **参数解释**： 辅助弹性网卡所属云服务实例的ID，例如RDS实例的ID，可以通过此字段过滤辅助弹性网卡，支持传入多个ID过滤。 **取值范围**： 不涉及。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	InstanceId *[]string `json:"instance_id,omitempty"`

	// **参数解释**： 辅助弹性网卡所属云服务实例的类型，例如“RDS”，可以通过此字段过滤辅助弹性网卡，支持传入多个实例类型过滤。 **取值范围**： 不涉及。 **约束限制**： 不涉及。 **默认取值**： 不涉及。
	InstanceType *[]string `json:"instance_type,omitempty"`
}

func (o ListSubNetworkInterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubNetworkInterfacesRequest struct{}"
	}

	return strings.Join([]string{"ListSubNetworkInterfacesRequest", string(data)}, " ")
}
