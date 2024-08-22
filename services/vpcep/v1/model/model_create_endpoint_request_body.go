package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateEndpointRequestBody 创建终端节接口请求结构体
type CreateEndpointRequestBody struct {

	// 创建连接Interface类型终端节点服务的终端节点时，此参数必选。 需要指定vpc_id对应VPC下已创建的网络（network）的ID，UUID格式。 说明： - VPC的子网网段不能与198.19.128.0/17重叠 - VPC路由表中自定义路由的目的地址不能与198.19.128.0/17重叠
	SubnetId *string `json:"subnet_id,omitempty"`

	// 终端节点服务的ID。 可以通过查询终端节点服务概 要获取要连接的终端节点服务 ID。
	EndpointServiceId string `json:"endpoint_service_id"`

	// 终端节点所在的VPC的ID。
	VpcId string `json:"vpc_id"`

	// 是否创建域名。  - true：创建域名  - false：不创建域名 默认值为false。 说明 当创建gateway类型终端节点服务的终端节点时， “enable_dns”设置为true或者false，均不创建域名。
	EnableDns *bool `json:"enable_dns,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]TagList `json:"tags,omitempty"`

	// 路由表ID列表。 创建gateway类型终端节点服务的终端节点时，此参数必选。 不设置此参数时，选择默认路由表。
	Routetables *[]string `json:"routetables,omitempty"`

	// 访问所连接的终端节点服务的IP。 创建终端节点时，可以指定访问所连接的终端节点服务的IP，目前只支持IPv4类型 。 创建连接Interface类型终端节点服务的终端节点时，此参数必选。
	PortIp *string `json:"port_ip,omitempty"`

	// 添加用于控制访问终端节点的白名单。 创建终端节点时，支持访问控制，使用此参数可以添加IPv4或CIDR，默认空列表。 仅当创建连接Interface类型终端节点服务的终端节点时，支持设置此参数。
	Whitelist *[]string `json:"whitelist,omitempty"`

	// 是否开启网络ACL隔离。
	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	// 描述字段，支持中英文字母、数字等字符，不支持“<”或“>”字符。
	Description *string `json:"description,omitempty"`

	// Gateway类型终端节点策略信息，仅限OBS、SFS的终端节点服务的enable_policy值为true时支持该参数。
	PolicyStatement *[]PolicyStatement `json:"policy_statement,omitempty"`

	// 终端节点策略信息，仅当终端节点服务的enable_policy值为true时支持该参数，默认值为完全访问权限。（OBS、SFS的终端节点服务暂不支持该参数）
	PolicyDocument *interface{} `json:"policy_document,omitempty"`

	// 指定终端节点的IP版本，仅专业型终端节点支持此参数。  - ipv4,  IPv4 - dualstack, 双栈
	IpVersion *CreateEndpointRequestBodyIpVersion `json:"ip_version,omitempty"`

	// 访问所连接的终端节点服务的IPv6的地址。  创建终端节点时，可以指定访问所连接的终端节点服务的IP，不指定的情况下，会使用系统生成的一个地址。  仅专业型终端节点支持此参数。
	Ipv6Address *string `json:"ipv6_address,omitempty"`
}

func (o CreateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequestBody", string(data)}, " ")
}

type CreateEndpointRequestBodyIpVersion struct {
	value string
}

type CreateEndpointRequestBodyIpVersionEnum struct {
	IPV4      CreateEndpointRequestBodyIpVersion
	DUALSTACK CreateEndpointRequestBodyIpVersion
}

func GetCreateEndpointRequestBodyIpVersionEnum() CreateEndpointRequestBodyIpVersionEnum {
	return CreateEndpointRequestBodyIpVersionEnum{
		IPV4: CreateEndpointRequestBodyIpVersion{
			value: "ipv4",
		},
		DUALSTACK: CreateEndpointRequestBodyIpVersion{
			value: "dualstack",
		},
	}
}

func (c CreateEndpointRequestBodyIpVersion) Value() string {
	return c.value
}

func (c CreateEndpointRequestBodyIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointRequestBodyIpVersion) UnmarshalJSON(b []byte) error {
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
