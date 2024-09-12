package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateEndpointServiceResponse Response Object
type UpdateEndpointServiceResponse struct {

	// 终端节点服务的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 标识终端节点服务后端资源的ID， 格式为通用唯一识别码（Universally Unique Identifier，下文简称UUID）。 取值为：  - LB类型：负载均衡器内网IP对应的端口ID。  - VM类型：弹性云服务器IP地址对应的网卡ID。  - VIP类型：虚拟资源所在物理服务器对应的网卡ID。（该字段已废弃，请优先使用LB类型）
	PortId *string `json:"port_id,omitempty"`

	// 终端节点服务的名称。
	ServiceName *string `json:"service_name,omitempty"`

	// 资源类型。  - VM：云服务器。  - VIP：虚拟IP。  - LB：增强负载均衡型。
	ServerType *string `json:"server_type,omitempty"`

	// 终端节点服务对应后端资源所在的VPC的ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 终端节点服务对应的集群id
	PoolId *string `json:"pool_id,omitempty"`

	// 是否需要审批。  - false：不需要审批，创建的终端节点连接直接为accepted状态。  - true：需要审批，创建的终端节点连接为pendingAcceptance状态， 需要终端节点服务所属用户审核后方可使用。
	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	// 终端节点服务的状态。  - creating：创建中  - available：可连接  - failed：失败
	Status *string `json:"status,omitempty"`

	// 终端节点服务类型。 终端节点服务类型包括“网关（gateway）型”和“接口（interface）型”：  - gateway：由运维人员配置。用户无需创建，可直接使用。  - interface：包括运维人员配置的云服务和用户自己创建的私有服务。 其中，运维人员配置的云服务无需创建，用户可直接使用。 您可以通过创建终端节点创建访问Gateway和Interface类型终端节点服务的终端节点。
	ServiceType *string `json:"service_type,omitempty"`

	// 终端节点服务的创建时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 终端节点服务的更新时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 接口型VLAN场景服务端IPv4地址或域名
	Ip *string `json:"ip,omitempty"`

	// 服务开放的端口映射列表 同一个终端节点服务下，不允许重复的端口映射。 若多个终端节点服务共用一个port_id， 则终端节点服务之间的所有端口映射的server_port和protocol的组合不能重复。
	Ports *[]PortList `json:"ports,omitempty"`

	// 用于控制将哪些信息（如客户端的源IP、源端口、marker_id等）携带到服务端。 支持携带的客户端信息包括如下两种类型：  - TCP TOA：表示将客户端信息插入到tcp option字段中携带至服务端。 说明：仅当后端资源为OBS时，支持TCP TOA类型信息携带方式。  - Proxy Protocol：表示将客户端信息插入到tcp payload字段中携带至服务端。 仅当服务端支持解析上述字段时，该参数设置才有效。 该参数的取值包括：  - close：表示关闭代理协议。  - toa_open：表示开启代理协议“tcp_toa”。  - proxy_open：表示开启代理协议“proxy_protocol”。  - open：表示同时开启代理协议“tcp_toa”和“proxy_protocol”。 默认值为“close”。
	TcpProxy *string `json:"tcp_proxy,omitempty"`

	// 资源标签列表
	Tags *[]TagList `json:"tags,omitempty"`

	// 描述字段，支持中英文字母、数字等字符，不支持“<”或“>”字符。
	Description *string `json:"description,omitempty"`

	// 是否允许自定义终端节点策略。  - false：不支持设置终端节点策略  - true：支持设置终端节点策略 默认为false
	EnablePolicy *bool `json:"enable_policy,omitempty"`

	// 指定终端节点服务的IP版本，仅专业型终端节点服务支持此参数 ● ipv4,  IPv4 ● ipv6,  IPv6
	IpVersion      *UpdateEndpointServiceResponseIpVersion `json:"ip_version,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o UpdateEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceResponse", string(data)}, " ")
}

type UpdateEndpointServiceResponseIpVersion struct {
	value string
}

type UpdateEndpointServiceResponseIpVersionEnum struct {
	IPV4 UpdateEndpointServiceResponseIpVersion
	IPV6 UpdateEndpointServiceResponseIpVersion
}

func GetUpdateEndpointServiceResponseIpVersionEnum() UpdateEndpointServiceResponseIpVersionEnum {
	return UpdateEndpointServiceResponseIpVersionEnum{
		IPV4: UpdateEndpointServiceResponseIpVersion{
			value: "ipv4",
		},
		IPV6: UpdateEndpointServiceResponseIpVersion{
			value: "ipv6",
		},
	}
}

func (c UpdateEndpointServiceResponseIpVersion) Value() string {
	return c.value
}

func (c UpdateEndpointServiceResponseIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointServiceResponseIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
