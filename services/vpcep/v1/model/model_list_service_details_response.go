package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ListServiceDetailsResponse struct {
	// 终端节点服务的ID，唯一标识。

	Id *string `json:"id,omitempty"`
	// 标识终端节点服务后端资源的ID，格式为通用唯一识别码 （Universally Unique Identifier，下文简称UUID）。取值为： ● LB类型：增强型负载均衡器内网IP对应的端口ID。 ● VM类型：弹性云服务器IP地址对应的网卡ID。 ● VIP类型：虚拟资源所在物理服务器对应的网卡ID。

	PortId *string `json:"port_id,omitempty"`
	// 虚拟IP的网卡ID。 仅当“port_id”为“VIP类型”时，返回该参数。

	VipPortId *string `json:"vip_port_id,omitempty"`
	// 终端节点服务的名称。

	ServiceName *string `json:"service_name,omitempty"`
	// 资源类型。 ● VM：云服务器。 ● VIP：虚拟IP。 ● LB：增强负载均衡型。

	ServiceType *string `json:"service_type,omitempty"`
	// 终端节点服务对应后端资源所在的VPC的ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 是否需要审批。 ● false：不需要审批，创建的终端节点连接直接为accepted状态。 ● true：需要审批，创建的终端节点连接为pendingAcceptance状态，需要终端节点服务所属用户审核后方可使用。

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`
	// 终端节点服务的状态。 ● creating：创建中 ● available：可连接 ● failed：失败 ● deleting：删除中

	Status *ListServiceDetailsResponseStatus `json:"status,omitempty"`
	// 终端节点服务类型。 终端节点服务类型包括“网关（gataway）型”和“接口（interface）型”： ● gataway：由运维人员配置。用户无需创建，可直接使用。 ● interface：包括运维人员配置的云服务和用户自己创建的私有服务。其中，运维人员配 置的云服务无需创建，用户可直接使用。 您可以通过创建终端节点创建访问Gateway和Interface类型终端节点服务的终端节点。

	ServerType *ListServiceDetailsResponseServerType `json:"server_type,omitempty"`
	// 终端节点服务的创建时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ

	CreatedAt *string `json:"created_at,omitempty"`
	// 终端节点服务的更新时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ

	UpdatedAt *string `json:"updated_at,omitempty"`
	// 项目ID，获取方法请参见获取项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 网段类型。 public：公网网段 internal：内网网段 默认值为internal。

	CidrType *string `json:"cidr_type,omitempty"`
	// 服务开放的端口映射列表，详细内容请参见表4-17 同一个终端节点服务下，不允许重复的端口映射。若多个终端节点服务共用一个port_id，则 终端节点服务之间的所有端口映射的server_port和protocol的组合不能重复。

	Ports *[]PortList `json:"ports,omitempty"`
	// 用于控制是否将客户端的源IP、源端口、marker_id等信息携带到服务端。信息携带支持两种方式： ● TCP TOA：表示将客户端信息插入到tcpoption字段中携带至服务端。 说明 仅当后端资源为OBS时，支持TCP TOA类型信息携带方式。 ● Proxy Protocol：表示将客户端相关信息插入到tcp payload字段中携带至服务端。 仅当服务端支持解析上述字段时，该参数设置才有效。 参数的取值包括： ● close：表示关闭代理协议。 ● toa_open：表示开启代理协议“tcp_toa”。 ● proxy_open：表示开启代理协议“proxy_protocol”。 ● open：表示同时开启代理协议“tcp_toa”和“proxy_protocol”。 默认值为“close”。

	TcpProxy *ListServiceDetailsResponseTcpProxy `json:"tcp_proxy,omitempty"`
	// 资源标签列表

	Tags *[]TagList `json:"tags,omitempty"`
	// 提交任务异常时返回的异常信息

	Error          *[]Error `json:"error,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListServiceDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceDetailsResponse", string(data)}, " ")
}

type ListServiceDetailsResponseStatus struct {
	value string
}

type ListServiceDetailsResponseStatusEnum struct {
	CREATING  ListServiceDetailsResponseStatus
	AVAILABLE ListServiceDetailsResponseStatus
	FAILED    ListServiceDetailsResponseStatus
}

func GetListServiceDetailsResponseStatusEnum() ListServiceDetailsResponseStatusEnum {
	return ListServiceDetailsResponseStatusEnum{
		CREATING: ListServiceDetailsResponseStatus{
			value: "creating",
		},
		AVAILABLE: ListServiceDetailsResponseStatus{
			value: "available",
		},
		FAILED: ListServiceDetailsResponseStatus{
			value: "failed",
		},
	}
}

func (c ListServiceDetailsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDetailsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ListServiceDetailsResponseServerType struct {
	value string
}

type ListServiceDetailsResponseServerTypeEnum struct {
	VM  ListServiceDetailsResponseServerType
	VIP ListServiceDetailsResponseServerType
	LB  ListServiceDetailsResponseServerType
}

func GetListServiceDetailsResponseServerTypeEnum() ListServiceDetailsResponseServerTypeEnum {
	return ListServiceDetailsResponseServerTypeEnum{
		VM: ListServiceDetailsResponseServerType{
			value: "VM",
		},
		VIP: ListServiceDetailsResponseServerType{
			value: "VIP",
		},
		LB: ListServiceDetailsResponseServerType{
			value: "LB",
		},
	}
}

func (c ListServiceDetailsResponseServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDetailsResponseServerType) UnmarshalJSON(b []byte) error {
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

type ListServiceDetailsResponseTcpProxy struct {
	value string
}

type ListServiceDetailsResponseTcpProxyEnum struct {
	CLOSE      ListServiceDetailsResponseTcpProxy
	TOA_OPEN   ListServiceDetailsResponseTcpProxy
	PROXY_OPEN ListServiceDetailsResponseTcpProxy
	OPEN       ListServiceDetailsResponseTcpProxy
}

func GetListServiceDetailsResponseTcpProxyEnum() ListServiceDetailsResponseTcpProxyEnum {
	return ListServiceDetailsResponseTcpProxyEnum{
		CLOSE: ListServiceDetailsResponseTcpProxy{
			value: "close",
		},
		TOA_OPEN: ListServiceDetailsResponseTcpProxy{
			value: "toa_open",
		},
		PROXY_OPEN: ListServiceDetailsResponseTcpProxy{
			value: "proxy_open",
		},
		OPEN: ListServiceDetailsResponseTcpProxy{
			value: "open",
		},
	}
}

func (c ListServiceDetailsResponseTcpProxy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDetailsResponseTcpProxy) UnmarshalJSON(b []byte) error {
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
