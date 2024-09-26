package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateEndpointServiceRequestBody 创建终端节点服务接口请求结构体
type CreateEndpointServiceRequestBody struct {

	// 标识终端节点服务后端资源的ID， 格式为通用唯一识别码（Universally Unique Identifier，下文简称UUID）。 取值为： - LB类型：负载均衡器内网IP对应的端口ID。 详细内容请参考《弹性负载均衡API参考》中的“查询负载均衡详情”。 - VM类型：弹性云服务器IP地址对应的网卡ID。 详细内容请参考《弹性云服务器API参考》中的“查询云服务器网卡信息”， 详见响应消息中的“port_id”字段。 - VIP类型：虚拟IP所在虚拟机的网卡ID（VIP类型业务已不支持，该取值类型已废弃） 说明： - 创建终端节点服务时，VPC的子网网段不能与198.19.128.0/17重叠。 - VPC路由表中自定义路由的目的地址不能与198.19.128.0/17重叠。
	PortId string `json:"port_id"`

	// 终端节点服务的名称，长度不大于16，允许传入大小写字母、数字、下划线、中划线。 - 传入为空，存入值为regionName+.+serviceId - 传入不为空并校验通过，存入值为regionName+.+serviceName+.+serviceId
	ServiceName *string `json:"service_name,omitempty"`

	// 终端节点服务对应后端资源所在的VPC的ID。
	VpcId string `json:"vpc_id"`

	// 是否需要审批。  - false：不需要审批，创建的终端节点连接直接为accepted状态。  - true：需要审批，创建的终端节点连接为pendingAcceptance状态， 需要终端节点服务所属用户审核后方可使用。 默认为true，需要审批。
	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	// 终端节点服务类型。 仅支持将用户私有服务创建为interface类型的终端节点服务。 终端节点服务类型包括“网关（gateway）型”和“接口（interface）型”：  - gateway：由运维人员配置。用户无需创建，可直接使用。  - interface：包括运维人员配置的云服务和用户自己创建的私有服务。 其中，运维人员配置的云服务无需创建， 用户可直接使用。 您可以通过查询公共终端节点服务列表， 查看由运维人员配置的所有用户可见且可连接的终端节点服务， 并通过创建终端节点创建访问Gateway和Interface类型终端节点服务的终端节点。
	ServiceType *CreateEndpointServiceRequestBodyServiceType `json:"service_type,omitempty"`

	// 资源类型。  - VM：云服务器，适用于作为服务器使用。  - VIP：虚拟IP，适用于作为虚IP场景使用。（该字段已废弃，请优先使用LB类型）  - LB：负载均衡，适用于高访问量业务和对可靠性和容灾性要求较高的业务。
	ServerType CreateEndpointServiceRequestBodyServerType `json:"server_type"`

	// 接口型VLAN场景服务端IPv4地址或域名
	Ip *string `json:"ip,omitempty"`

	// 服务开放的端口映射列表，详细内容请参见表4-10。 同一个终端节点服务下，不允许重复的端口映射。若多个终端节点服务共用一个port_id， 则终端节点服务之间的所有端口映射的server_port和protocol的组合不能重复， 单次最多添加200个。
	Ports []PortList `json:"ports"`

	// 用于控制将哪些信息（如客户端的源IP、源端口、marker_id等）携带到服务端。 支持携带的客户端信息包括如下两种类型：  - TCP TOA：表示将客户端信息插入到tcp option字段中携带至服务端。 说明：仅当后端资源为OBS时，支持TCP TOA类型信息携带方式。  - Proxy Protocol：表示将客户端信息插入到tcp payload字段中携带至服务端。 仅当服务端支持解析上述字段时，该参数设置才有效。 该参数的取值包括：  - close：表示关闭代理协议。  - toa_open：表示开启代理协议“tcp_toa”。  - proxy_open：表示开启代理协议“proxy_protocol”。  - open：表示同时开启代理协议“tcp_toa”和“proxy_protocol”。 默认值为“close”。
	TcpProxy *CreateEndpointServiceRequestBodyTcpProxy `json:"tcp_proxy,omitempty"`

	// 资源标签列表。同一个终端节点服务最多可添加20个标签。
	Tags *[]TagList `json:"tags,omitempty"`

	// 描述字段，支持中英文字母、数字等字符，不支持“<”或“>”字符。  描述字段，支持中英文字母、数字等字符，不支持“<”或“>”字符。
	Description *string `json:"description,omitempty"`

	// 指定终端节点服务的IP版本，仅专业型终端节点服务支持此参数 ● ipv4,  IPv4 ● ipv6,  IPv6
	IpVersion *CreateEndpointServiceRequestBodyIpVersion `json:"ip_version,omitempty"`

	// 接口型snat的地址段，ip_version为ipv6时必选。创建服务时使用的VPC内的任意一个网络ID。当服务类型为VIP、VM、ELBV2类型时使用
	SnatNetworkId *string `json:"snat_network_id,omitempty"`
}

func (o CreateEndpointServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointServiceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointServiceRequestBody", string(data)}, " ")
}

type CreateEndpointServiceRequestBodyServiceType struct {
	value string
}

type CreateEndpointServiceRequestBodyServiceTypeEnum struct {
	GATEWAY   CreateEndpointServiceRequestBodyServiceType
	INTERFACE CreateEndpointServiceRequestBodyServiceType
}

func GetCreateEndpointServiceRequestBodyServiceTypeEnum() CreateEndpointServiceRequestBodyServiceTypeEnum {
	return CreateEndpointServiceRequestBodyServiceTypeEnum{
		GATEWAY: CreateEndpointServiceRequestBodyServiceType{
			value: "gateway",
		},
		INTERFACE: CreateEndpointServiceRequestBodyServiceType{
			value: "interface",
		},
	}
}

func (c CreateEndpointServiceRequestBodyServiceType) Value() string {
	return c.value
}

func (c CreateEndpointServiceRequestBodyServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyServiceType) UnmarshalJSON(b []byte) error {
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

type CreateEndpointServiceRequestBodyServerType struct {
	value string
}

type CreateEndpointServiceRequestBodyServerTypeEnum struct {
	VM  CreateEndpointServiceRequestBodyServerType
	VIP CreateEndpointServiceRequestBodyServerType
	LB  CreateEndpointServiceRequestBodyServerType
}

func GetCreateEndpointServiceRequestBodyServerTypeEnum() CreateEndpointServiceRequestBodyServerTypeEnum {
	return CreateEndpointServiceRequestBodyServerTypeEnum{
		VM: CreateEndpointServiceRequestBodyServerType{
			value: "VM",
		},
		VIP: CreateEndpointServiceRequestBodyServerType{
			value: "VIP",
		},
		LB: CreateEndpointServiceRequestBodyServerType{
			value: "LB",
		},
	}
}

func (c CreateEndpointServiceRequestBodyServerType) Value() string {
	return c.value
}

func (c CreateEndpointServiceRequestBodyServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyServerType) UnmarshalJSON(b []byte) error {
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

type CreateEndpointServiceRequestBodyTcpProxy struct {
	value string
}

type CreateEndpointServiceRequestBodyTcpProxyEnum struct {
	CLOSE      CreateEndpointServiceRequestBodyTcpProxy
	TOA_OPEN   CreateEndpointServiceRequestBodyTcpProxy
	PROXY_OPEN CreateEndpointServiceRequestBodyTcpProxy
	OPEN       CreateEndpointServiceRequestBodyTcpProxy
	PROXY_VNI  CreateEndpointServiceRequestBodyTcpProxy
}

func GetCreateEndpointServiceRequestBodyTcpProxyEnum() CreateEndpointServiceRequestBodyTcpProxyEnum {
	return CreateEndpointServiceRequestBodyTcpProxyEnum{
		CLOSE: CreateEndpointServiceRequestBodyTcpProxy{
			value: "close",
		},
		TOA_OPEN: CreateEndpointServiceRequestBodyTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: CreateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_open",
		},
		OPEN: CreateEndpointServiceRequestBodyTcpProxy{
			value: "open",
		},
		PROXY_VNI: CreateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c CreateEndpointServiceRequestBodyTcpProxy) Value() string {
	return c.value
}

func (c CreateEndpointServiceRequestBodyTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyTcpProxy) UnmarshalJSON(b []byte) error {
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

type CreateEndpointServiceRequestBodyIpVersion struct {
	value string
}

type CreateEndpointServiceRequestBodyIpVersionEnum struct {
	IPV4 CreateEndpointServiceRequestBodyIpVersion
	IPV6 CreateEndpointServiceRequestBodyIpVersion
}

func GetCreateEndpointServiceRequestBodyIpVersionEnum() CreateEndpointServiceRequestBodyIpVersionEnum {
	return CreateEndpointServiceRequestBodyIpVersionEnum{
		IPV4: CreateEndpointServiceRequestBodyIpVersion{
			value: "ipv4",
		},
		IPV6: CreateEndpointServiceRequestBodyIpVersion{
			value: "ipv6",
		},
	}
}

func (c CreateEndpointServiceRequestBodyIpVersion) Value() string {
	return c.value
}

func (c CreateEndpointServiceRequestBodyIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyIpVersion) UnmarshalJSON(b []byte) error {
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
