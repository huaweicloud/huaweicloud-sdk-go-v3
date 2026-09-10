package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubNetworkInterface struct {

	// **参数解释**： 辅助弹性网卡的资源ID。辅助弹性网卡创建成功后，会生成一个辅助弹性网卡 ID，是辅助弹性网卡对应的唯一标识。 **取值范围**： 带“-”的标准UUID格式。
	Id string `json:"id"`

	// **参数解释**： 辅助弹性网卡所在的虚拟子网ID。 **取值范围**： 带“-”的标准UUID格式。
	VirsubnetId string `json:"virsubnet_id"`

	// **参数解释**： 辅助弹性网卡的私有IPv4地址。 **取值范围**： 不涉及。
	PrivateIpAddress string `json:"private_ip_address"`

	// **参数解释**： 辅助弹性网卡的私有IPv6地址。 **取值范围**： 不涉及。
	Ipv6IpAddress string `json:"ipv6_ip_address"`

	// **参数解释**： 辅助弹性网卡的MAC地址。 **取值范围**： 合法的MAC地址，系统随机分配。
	MacAddress string `json:"mac_address"`

	// **参数解释**： 辅助弹性网卡的宿主网卡所属的设备ID。 **取值范围**： 带“-”的标准UUID格式。
	ParentDeviceId string `json:"parent_device_id"`

	// **参数解释**： 辅助弹性网卡所挂载的弹性网卡的ID。 **取值范围**： 带“-”的标准UUID格式。
	ParentId string `json:"parent_id"`

	// **参数解释**： 辅助弹性网卡的描述信息。 **取值范围**： 0-255个字符，不能包含“<”和“>”。
	Description string `json:"description"`

	// **参数解释**： 辅助弹性网卡所属VPC的ID。 **取值范围**： 带“-”的标准UUID格式。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 辅助弹性网卡的VLAN ID。 **取值范围**： 1-4094
	VlanId int32 `json:"vlan_id"`

	// **参数解释**： 辅助弹性网卡关联的安全组的ID列表。例如：\"security_groups\": [\"a0608cbf-d047-4f54-8b28-cd7b59853fff\"]。 **取值范围**： 如果请求时不指定此参数，辅助弹性网卡创建后会自动关联默认安全组。
	SecurityGroups []string `json:"security_groups"`

	// **参数解释**： 辅助弹性网卡的标签信息，包括标签键和标签值，可用来分类和标识资源。详情请参见Tag对象。 **取值范围**： 不涉及。
	Tags []ResponseTag `json:"tags"`

	// **参数解释**： 辅助弹性网卡所属的项目ID。 **取值范围**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 辅助弹性网卡的创建时间。 **取值范围**： 不涉及。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// **参数解释**： 辅助弹性网卡的更新时间。 **取值范围**： 不涉及。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// **参数解释**： 辅助弹性网卡的IP/Mac对列表，详情请参见“AllowedAddressPair”对象表。 **取值范围**： 不涉及。
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs"`

	// **参数解释**： 辅助弹性网卡的状态。 **取值范围**： - NORMAL：表示辅助弹性网卡已挂载在弹性网卡上。 - UNBOUND：表示辅助弹性网卡未挂载在弹性网卡上。
	State string `json:"state"`

	// **参数解释**： 辅助弹性网卡所属的云服务实例ID，例如RDS实例ID。 **取值范围**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 辅助弹性网卡所属的云服务实例类型，例如“RDS”。 **取值范围**： 不涉及。
	InstanceType string `json:"instance_type"`

	// **参数解释**： 辅助弹性网卡所在站点的公网出口信息。 **取值范围**： - center：默认值，表示作用域为中心。 - 某个AZ ID：表示作用域为具体的AZ。
	Scope string `json:"scope"`

	// **参数解释**： 辅助弹性网卡安全使能标记，如果不使能则安全组不生效。 **取值范围**： 不涉及。
	SecurityEnabled bool `json:"security_enabled"`
}

func (o SubNetworkInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubNetworkInterface struct{}"
	}

	return strings.Join([]string{"SubNetworkInterface", string(data)}, " ")
}
