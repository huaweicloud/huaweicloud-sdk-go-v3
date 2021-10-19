package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

type ServiceDetailsResp struct {
	// 终端节点服务的ID，唯一标识。

	Id *string `json:"id,omitempty"`
	// 标识终端节点服务后端资源的ID，格式为通用 唯一识别码（Universally Unique Identifier， 下文简称UUID）。取值为： ● LB类型：增强型负载均衡器内网IP对应的端 口ID。 ● VM类型：弹性云服务器IP地址对应的网卡 ID。 ● VIP类型：虚拟资源所在物理服务器对应的 网卡ID。

	PortId *string `json:"port_id,omitempty"`
	// 虚拟IP的网卡ID。 仅当“port_id”为“VIP类型”时，返回该参 数。

	VipPortId *string `json:"vip_port_id,omitempty"`
	// 终端节点服务的名称。

	ServiceName *string `json:"service_name,omitempty"`
	// 资源类型。 ● VM：云服务器。 ● VIP：虚拟IP。 ● LB：增强负载均衡型。

	ServiceType *string `json:"service_type,omitempty"`
	// 终端节点服务对应后端资源所在的VPC的ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 是否需要审批。 ● false：不需要审批，创建的终端节点连接直 接为accepted状态。 ● true：需要审批，创建的终端节点连接为 pendingAcceptance状态，需要终端节点服 务所属用户审核后方可使用。

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`
	// 终端节点服务的状态。 ● creating：创建中 ● available：可连接 ● failed：失败 ● deleting：删除中

	Status *ServiceDetailsRespStatus `json:"status,omitempty"`
	// 终端节点服务类型。 终端节点服务类型包括“网关（gataway） 型”和“接口（interface）型”： ● gataway：由运维人员配置。用户无需创 建，可直接使用。 ● interface：包括运维人员配置的云服务和用 户自己创建的私有服务。其中，运维人员配 置的云服务无需创建，用户可直接使用。 您可以通过创建终端节点创建访问Gateway和 Interface类型终端节点服务的终端节点。

	ServerType *ServiceDetailsRespServerType `json:"server_type,omitempty"`
	// 终端节点服务的创建时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH: MM:SSZ

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 终端节点服务的更新时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH: MM:SSZ

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
	// 项目ID，获取方法请参见获取项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 网段类型。 ● public：公网网段 ● internal：内网网段 默认值为internal。

	CidrType *ServiceDetailsRespCidrType `json:"cidr_type,omitempty"`
	// 服务开放的端口映射列表，详细内容请参见表 4-17 同一个终端节点服务下，不允许重复的端口映 射。若多个终端节点服务共用一个port_id，则 终端节点服务之间的所有端口映射的 server_port和protocol的组合不能重复。

	Ports *[]PortList `json:"ports,omitempty"`
	// 用于控制是否将客户端的源IP、源端口、 marker_id等信息携带到服务端。信息携带支 持两种方式： ● TCP TOA：表示将客户端信息插入到tcp option字段中携带至服务端。 说明 仅当后端资源为OBS时，支持TCP TOA类型信息 携带方式。 ● Proxy Protocol：表示将客户端相关信息插 入到tcp payload字段中携带至服务端。 仅当服务端支持解析上述字段时，该参数设置 才有效。 参数的取值包括： ● close：表示关闭代理协议。 ● toa_open：表示开启代理协议 “tcp_toa”。 ● proxy_open：表示开启代理协议 “proxy_protocol”。 ● open：表示同时开启代理协议“tcp_toa” 和“proxy_protocol”。 默认值为“close”。

	TcpProxy *ServiceDetailsRespTcpProxy `json:"tcp_proxy,omitempty"`
	// 资源标签列表

	Tags *[]TagList `json:"tags,omitempty"`
	// 错误信息。 当终端节点服务状态异常，即“status”的值 为“failed”时，会返回该字段，详细内容请 参见表4-19。

	Error *[]QueryError `json:"error,omitempty"`
}

func (o ServiceDetailsResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceDetailsResp struct{}"
	}

	return strings.Join([]string{"ServiceDetailsResp", string(data)}, " ")
}

type ServiceDetailsRespStatus struct {
	value string
}

type ServiceDetailsRespStatusEnum struct {
	CREATING  ServiceDetailsRespStatus
	AVAILABLE ServiceDetailsRespStatus
	FAILED    ServiceDetailsRespStatus
}

func GetServiceDetailsRespStatusEnum() ServiceDetailsRespStatusEnum {
	return ServiceDetailsRespStatusEnum{
		CREATING: ServiceDetailsRespStatus{
			value: "creating",
		},
		AVAILABLE: ServiceDetailsRespStatus{
			value: "available",
		},
		FAILED: ServiceDetailsRespStatus{
			value: "failed",
		},
	}
}

func (c ServiceDetailsRespStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ServiceDetailsRespStatus) UnmarshalJSON(b []byte) error {
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

type ServiceDetailsRespServerType struct {
	value string
}

type ServiceDetailsRespServerTypeEnum struct {
	VM  ServiceDetailsRespServerType
	VIP ServiceDetailsRespServerType
	LB  ServiceDetailsRespServerType
}

func GetServiceDetailsRespServerTypeEnum() ServiceDetailsRespServerTypeEnum {
	return ServiceDetailsRespServerTypeEnum{
		VM: ServiceDetailsRespServerType{
			value: "VM",
		},
		VIP: ServiceDetailsRespServerType{
			value: "VIP",
		},
		LB: ServiceDetailsRespServerType{
			value: "LB",
		},
	}
}

func (c ServiceDetailsRespServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ServiceDetailsRespServerType) UnmarshalJSON(b []byte) error {
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

type ServiceDetailsRespCidrType struct {
	value string
}

type ServiceDetailsRespCidrTypeEnum struct {
	PUBLIC   ServiceDetailsRespCidrType
	INTERNAL ServiceDetailsRespCidrType
}

func GetServiceDetailsRespCidrTypeEnum() ServiceDetailsRespCidrTypeEnum {
	return ServiceDetailsRespCidrTypeEnum{
		PUBLIC: ServiceDetailsRespCidrType{
			value: "public",
		},
		INTERNAL: ServiceDetailsRespCidrType{
			value: "internal",
		},
	}
}

func (c ServiceDetailsRespCidrType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ServiceDetailsRespCidrType) UnmarshalJSON(b []byte) error {
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

type ServiceDetailsRespTcpProxy struct {
	value string
}

type ServiceDetailsRespTcpProxyEnum struct {
	CLOSE      ServiceDetailsRespTcpProxy
	TOA_OPEN   ServiceDetailsRespTcpProxy
	PROXY_OPEN ServiceDetailsRespTcpProxy
	OPEN       ServiceDetailsRespTcpProxy
}

func GetServiceDetailsRespTcpProxyEnum() ServiceDetailsRespTcpProxyEnum {
	return ServiceDetailsRespTcpProxyEnum{
		CLOSE: ServiceDetailsRespTcpProxy{
			value: "close",
		},
		TOA_OPEN: ServiceDetailsRespTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: ServiceDetailsRespTcpProxy{
			value: "proxy_open",
		},
		OPEN: ServiceDetailsRespTcpProxy{
			value: "open",
		},
	}
}

func (c ServiceDetailsRespTcpProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ServiceDetailsRespTcpProxy) UnmarshalJSON(b []byte) error {
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
