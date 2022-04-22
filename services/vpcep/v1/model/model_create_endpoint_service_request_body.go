package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建终端节点服务接口请求结构体
type CreateEndpointServiceRequestBody struct {

	// 标识终端节点服务后端资源的ID，格式为通用唯一识别码 （Universally Unique Identifier，下文简称UUID）。 取值为： ● LB类型：增强型负载均衡器内网IP对应的端口ID。详细内容请参考《弹性负载均衡API参考》中的“查询负载均 衡详情”，详见响应消息中的“vip_port_id”字段。 ● VM类型：弹性云服务器IP地址对应的网卡ID。详细内容请参考《弹性云服务器API参考》中的“查询云服务器网 卡信息”，详见响应消息中的“port_id”字段。 ● VIP类型：虚拟资源所在物理服务器对应的网卡ID。 说明 ● 创建终端节点服务时，VPC的子网网段不能与198.19.128.0/17重叠。 ● VPC路由表中自定义路由的目的地址不能与198.19.128.0/17重叠
	PortId string `json:"port_id"`

	// 虚拟IP的网卡ID。
	VipPortId *string `json:"vip_port_id,omitempty"`

	// 终端节点服务的名称，长度不大于16，允许传入大小写字母、数字、下划线、中划线。 ● 传入为空，存入值为regionName+.+serviceId ● 传入不为空并校验通过，存入值为regionName+.+serviceName+.+serviceId
	ServiceName *string `json:"service_name,omitempty"`

	// 终端节点服务对应后端资源所在的VPC的ID。 详细内容请参考《虚拟私有云API参考》中的“查询VPC”，详见响应消息中的“id”字段。
	VpcId string `json:"vpc_id"`

	// 是否需要审批。 ● false：不需要审批，创建的终端节点连接直接为accepted状态。 ● true：需要审批，创建的终端节点连接为pendingAcceptance状态，需要终端节点服务所属用户审核后方可使用。 默认为true，需要审批。
	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	// 终端节点服务类型。仅支持将用户私有服务创建为interface类型的终端节点服务。 终端节点服务类型包括“网关（gataway）型”和“接口（interface）型”： ● gataway：由运维人员配置。用户无需创建，可直接使用。 ● interface：包括运维人员配置的云服务和用户自己创建的私有服务。其中，运维人员配置的云服务无需创建， 用户可直接使用。 您可以通过查询公共终端节点服务列表查看由运维人员配置的所有用户可见且可连接的终端节点服务，并通过创建终 端节点创建访问Gateway和Interface类型终端节点服务的终端节点。
	ServiceType *string `json:"service_type,omitempty"`

	// 资源类型。 ● VM：云服务器，适用于作为服务器使用。 ● VIP：虚拟IP，适用于作为虚拟资源的物理服务器使用。 ● LB：增强型负载均衡，适用于高访问量业务和对可靠性和容灾性要求较高的业务。
	ServerType CreateEndpointServiceRequestBodyServerType `json:"server_type"`

	// 服务开放的端口映射列表，详细内容请参见表4-10。 同一个终端节点服务下，不允许重复的端口映射。若多个终端节点服务共用一个port_id， 则终端节点服务之间的所有端口映射的server_port和protocol的组合不能重复，单次最多添加200个。
	Ports []PortList `json:"ports"`

	// 用于控制是否将客户端的源IP、源端口、marker_id等信息携带到服务端。信息携带支持两种方式： ● TCP TOA：表示将客户端信息插入到tcp option字段中携带至服务端。 说明 仅当后端资源为OBS时，支持TCP TOA类型信息携带方式。 ● Proxy Protocol：表示将客户端相关信息插入到tcp payload字段中携带至服务端。 仅当服务端支持解析上述字段时，该参数设置才有效。 参数的取值包括： ● close：表示关闭代理协议。 ● toa_open：表示开启代理协议“tcp_toa”。 ● proxy_open：表示开启代理协议“proxy_protocol”。 ● open：表示同时开启代理协议“tcp_toa”和“proxy_protocol”。 默认值为“close”。
	TcpProxy *CreateEndpointServiceRequestBodyTcpProxy `json:"tcp_proxy,omitempty"`

	// 资源标签列表。同一个终端节点服务最多可添加10个标签。
	Tags *[]TagList `json:"tags,omitempty"`
}

func (o CreateEndpointServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointServiceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointServiceRequestBody", string(data)}, " ")
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

func (c CreateEndpointServiceRequestBodyServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyServerType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
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
	}
}

func (c CreateEndpointServiceRequestBodyTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointServiceRequestBodyTcpProxy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
