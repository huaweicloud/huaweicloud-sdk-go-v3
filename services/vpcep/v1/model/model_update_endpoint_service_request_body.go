package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEndpointServiceRequestBody 修改终端节点服务接口请求结构体
type UpdateEndpointServiceRequestBody struct {

	// 是否需要审批。 ● false：不需审批，创建的终端节点连接直接为accepted状态。 ● true：需审批，创建的终端节点连接需要终端节点服务所属用户审核后方可使用。 默认为true，需要审批。
	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`

	// 终端节点服务的名称，长度不大于16，允许传入大小写字母、数字、下划线、中划线。
	ServiceName *string `json:"service_name,omitempty"`

	// 服务开放的端口映射列表，同一个终端节点服务下，不允许重复的端口映射。 若多个终端节点服务共用一个port_id， 则终端节点之间服务的所有端口映射的server_port和protocol的组合不能重复， 单次最多添加200个。 该参数值将被全量更新。
	Ports *[]PortList `json:"ports,omitempty"`

	// 标识终端节点服务后端资源的ID， 格式为通用唯一识别码（Universally UniqueIdentifier，下文简称UUID）。 取值为： ● LB类型：负载均衡器内网IP对应的端口ID。 详细内容请参考《弹性负载均衡API参考》中的“查询负载均衡详情”。 ● VM类型：弹性云服务器IP地址对应的网卡ID。 详细内容请参考《弹性云服务器API参考》中的“查询云服务器网卡信息”， 详见响应消息中的“port_id”字段。 ● VIP类型：虚拟资源所在物理服务器对应的网卡ID。（该字段已废弃，请优先使用LB类型） 说明： 当后端资源为“LB类型”时，仅支持修改为同类型后端资源。 例如，共享型负载均衡仅支持更换为共享型负载均衡，不支持更换为独享型负载均衡。
	PortId *string `json:"port_id,omitempty"`

	// 用于控制将哪些信息（如客户端的源IP、源端口、marker_id等）携带到服务端。 支持携带的客户端信息包括如下两种类型： ● TCP TOA：表示将客户端信息插入到tcp option字段中携带至服务端。 说明：仅当后端资源为OBS时，支持TCP TOA类型信息携带方式。 ● Proxy Protocol：表示将客户端信息插入到tcp payload字段中携带至服务端。 仅当服务端支持解析上述字段时，该参数设置才有效。 该参数的取值包括： ● close：表示关闭代理协议。 ● toa_open：表示开启代理协议“tcp_toa”。 ● proxy_open：表示开启代理协议“proxy_protocol”。 ● open：表示同时开启代理协议“tcp_toa”和“proxy_protocol”。 ● proxy_vni: 关闭toa，开启proxy和vni。 默认值为“close”。
	TcpProxy *UpdateEndpointServiceRequestBodyTcpProxy `json:"tcp_proxy,omitempty"`

	// 描述字段，支持中英文字母、数字等字符，不支持“<”或“>”字符。
	Description *string `json:"description,omitempty"`
}

func (o UpdateEndpointServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceRequestBody", string(data)}, " ")
}

type UpdateEndpointServiceRequestBodyTcpProxy struct {
	value string
}

type UpdateEndpointServiceRequestBodyTcpProxyEnum struct {
	CLOSE      UpdateEndpointServiceRequestBodyTcpProxy
	TOA_OPEN   UpdateEndpointServiceRequestBodyTcpProxy
	PROXY_OPEN UpdateEndpointServiceRequestBodyTcpProxy
	OPEN       UpdateEndpointServiceRequestBodyTcpProxy
	PROXY_VNI  UpdateEndpointServiceRequestBodyTcpProxy
}

func GetUpdateEndpointServiceRequestBodyTcpProxyEnum() UpdateEndpointServiceRequestBodyTcpProxyEnum {
	return UpdateEndpointServiceRequestBodyTcpProxyEnum{
		CLOSE: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "close",
		},
		TOA_OPEN: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_open",
		},
		OPEN: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "open",
		},
		PROXY_VNI: UpdateEndpointServiceRequestBodyTcpProxy{
			value: "proxy_vni",
		},
	}
}

func (c UpdateEndpointServiceRequestBodyTcpProxy) Value() string {
	return c.value
}

func (c UpdateEndpointServiceRequestBodyTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointServiceRequestBodyTcpProxy) UnmarshalJSON(b []byte) error {
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
