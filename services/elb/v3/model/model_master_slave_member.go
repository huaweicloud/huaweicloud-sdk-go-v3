package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasterSlaveMember **参数解释**：后端服务器信息。
type MasterSlaveMember struct {

	// **参数解释**：后端服务器ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：后端服务器名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：后端服务器的管理状态。 虽然创建、更新请求支持该字段，但实际取值决定于后端服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。  **取值范围**：true、false。
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**：后端服务器所在子网的IPv4子网ID或IPv6子网ID。 若所属的LB的IP类型后端转发特性已开启，则该字段可以不传，表示添加IP类型的后端服务器。此时address必须为**私网IPv4**地址，所在的pool的协议必须为UDP/TCP/TLS/HTTP/HTTPS/QUIC/GRPC[/GRPCS](tag:not_open)。  **取值范围**：不涉及  [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt)
	SubnetCidrId string `json:"subnet_cidr_id"`

	// **参数解释**：后端服务器业务端口。  >在开启端口透传的pool下创建member传该字段不生效，可不传该字段。
	ProtocolPort int32 `json:"protocol_port"`

	// **参数解释**：后端服务器对应的IP地址。  **取值范围**：不涉及  [不支持IPv6，请勿设置为IPv6地址。](tag:dt)
	Address string `json:"address"`

	// **参数解释**：当前后端服务器的IP地址版本，由后端系统自动根据传入的address字段确定。  **取值范围**：v4、v6
	IpVersion string `json:"ip_version"`

	// **参数解释**：设备所有者。  **取值范围**：不涉及 - 空，表示后端服务器未关联到ECS。 - compute:{az_name}，表示关联到ECS，其中{az_name}表示ECS所在可用区名。 - compute:subeni，表示辅助弹性网卡。  不支持该字段，请勿使用。
	DeviceOwner string `json:"device_owner"`

	// **参数解释**：关联的ECS ID，为空表示后端服务器未关联到ECS。  **取值范围**：不涉及  不支持该字段，请勿使用。
	DeviceId string `json:"device_id"`

	// **参数解释**：后端服务器的健康状态。  **取值范围**：不涉及 - ONLINE：后端服务器正常。 - NO_MONITOR：后端服务器所在的服务器组没有健康检查器。 - OFFLINE：后端服务器关联的ECS服务器不存在或已关机。 - INITIAL：后端云服务器健康检查打开时的初始状态。 - UNKNOWN: 后端云服务器组没有绑定监听器或者后端云服务器没有关联ECS等原因，后端云服务器健康检查结果未知。
	OperatingStatus string `json:"operating_status"`

	Reason *MemberHealthCheckFailedReason `json:"reason,omitempty"`

	// **参数解释**：后端服务器的类型。  **取值范围**： - ip：IP类型的member。 - instance：关联到ECS的member。
	MemberType string `json:"member_type"`

	// **参数解释**：member关联的实例ID。空表示member关联的实例为非真实设备 （如：IP类型的后端场景）。  **取值范围**：不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**：后端服务器的主备状态。  **取值范围**：不涉及
	Role string `json:"role"`

	// **参数解释**：后端服务器监听器粒度的的健康状态。 若绑定的监听器在该字段中，则以该字段中监听器对应的operating_status为准。 若绑定的监听器不在该字段中，则以外层的operating_status为准。  **取值范围**：不涉及
	Status []ListenerMemberInfo `json:"status"`
}

func (o MasterSlaveMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterSlaveMember struct{}"
	}

	return strings.Join([]string{"MasterSlaveMember", string(data)}, " ")
}
