package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateEndpointResponse Response Object
type CreateEndpointResponse struct {

	// 终端节点的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 终端节点连接的终端节点服务类型。 ● gataway：由运维人员配置。用户无需创建，可直接使用。 ● interface：包括运维人员配置的云服务和用户自己创建的私有服务。 其中，运维人员配置的云服务无需创建，用户可直接使用。 您可以通过查询公共终端节点服务列表， 查看由运维人员配置的所有用户可见且可连接的终端节点服务， 并通过创建终端节点服务创建Interface类型的终端节点服务。
	ServiceType *CreateEndpointResponseServiceType `json:"service_type,omitempty"`

	// 终端节点的状态。 ● pendingAcceptance：待接受 ● creating：创建中 ● accepted：已接受 ● rejected：已拒绝 ● failed：失败 ● deleting：删除中
	Status *CreateEndpointResponseStatus `json:"status,omitempty"`

	// 终端节点ip
	Ip *string `json:"ip,omitempty"`

	// 帐号状态。 ● frozen：冻结 ● active：解冻
	ActiveStatus *[]string `json:"active_status,omitempty"`

	// 终端节点服务的名称。
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	// 终端节点的报文标识。
	MarkerId *int32 `json:"marker_id,omitempty"`

	// 终端节点服务的ID。
	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	// 是否创建域名。 ● true：创建域名 ● false：不创建域名 说明 当创建连接gateway类型终端节点服务的终端节点时， “enable_dns”设置为true或者false，均不创建域名。
	EnableDns *bool `json:"enable_dns,omitempty"`

	// 访问所连接的终端节点服务的域名。 当“enable_dns”为true时，该参数可见。
	DnsNames *[]string `json:"dns_names,omitempty"`

	// vpc_id对应VPC下已创建的网络（network）的ID，UUID格式。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 终端节点所在的VPC的ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 终端节点的创建时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`

	// 终端节点的更新时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 项目ID，获取方法请参见获取项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]TagList `json:"tags,omitempty"`

	// 控制访问终端节点的白名单。 若未创建，则返回空列表。 创建连接Interface类型终端节点服务的终端节点时，显示此参数。
	Whitelist *[]string `json:"whitelist,omitempty"`

	// 是否开启网络ACL隔离。 ● true：开启网络ACL隔离 ● false：不开启网络ACL隔离 若未指定，则返回false。 创建连接Interface类型终端节点服务的终端节点时，显示此参数。
	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	// 路由表ID列表。 若未指定，返回默认VPC下路由表ID。 创建连接Gateway类型终端节点服务的终端节点时，显示此参数。
	Routetables *[]string `json:"routetables,omitempty"`

	// 规格名称
	SpecificationName *string `json:"specification_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 只涉及开启双端固定的网关型终端节点，响应体展示此字段
	PolicyStatement *[]string `json:"policy_statement,omitempty"`

	// 终端节点是否已停用，取值【enable/disable】
	EnableStatus *string `json:"enable_status,omitempty"`

	// 待废弃，实例相关联的集群ID
	EndpointPoolId *string `json:"endpoint_pool_id,omitempty"`

	// 终端节点对应Pool的Public Border Group信息
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o CreateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointResponse", string(data)}, " ")
}

type CreateEndpointResponseServiceType struct {
	value string
}

type CreateEndpointResponseServiceTypeEnum struct {
	INTERFACE CreateEndpointResponseServiceType
	GATEWAY   CreateEndpointResponseServiceType
}

func GetCreateEndpointResponseServiceTypeEnum() CreateEndpointResponseServiceTypeEnum {
	return CreateEndpointResponseServiceTypeEnum{
		INTERFACE: CreateEndpointResponseServiceType{
			value: "interface",
		},
		GATEWAY: CreateEndpointResponseServiceType{
			value: "gateway",
		},
	}
}

func (c CreateEndpointResponseServiceType) Value() string {
	return c.value
}

func (c CreateEndpointResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointResponseServiceType) UnmarshalJSON(b []byte) error {
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

type CreateEndpointResponseStatus struct {
	value string
}

type CreateEndpointResponseStatusEnum struct {
	PENDING_ACCEPTANCE CreateEndpointResponseStatus
	CREATING           CreateEndpointResponseStatus
	ACCEPTED           CreateEndpointResponseStatus
	REJECTED           CreateEndpointResponseStatus
	FAILED             CreateEndpointResponseStatus
	DELETING           CreateEndpointResponseStatus
}

func GetCreateEndpointResponseStatusEnum() CreateEndpointResponseStatusEnum {
	return CreateEndpointResponseStatusEnum{
		PENDING_ACCEPTANCE: CreateEndpointResponseStatus{
			value: "pendingAcceptance",
		},
		CREATING: CreateEndpointResponseStatus{
			value: "creating",
		},
		ACCEPTED: CreateEndpointResponseStatus{
			value: "accepted",
		},
		REJECTED: CreateEndpointResponseStatus{
			value: "rejected",
		},
		FAILED: CreateEndpointResponseStatus{
			value: "failed",
		},
		DELETING: CreateEndpointResponseStatus{
			value: "deleting",
		},
	}
}

func (c CreateEndpointResponseStatus) Value() string {
	return c.value
}

func (c CreateEndpointResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointResponseStatus) UnmarshalJSON(b []byte) error {
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
