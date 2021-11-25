package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建终端节接口请求结构体
type CreateEndpointRequestBody struct {
	// 说明： 创建Interface类型Client必选。需要指定vpc_id对应VPC下已 创建的网络（network）的 ID，UUID格式。 详细内容请参考《虚拟私有云 API参考》中的“查询子 网”，详见响应消息中的 “id”字段。 创建连接Interface类型终端节 点服务的终端节点时，此参数 必选。 说明 ● VPC的子网网段不能与 198.19.128.0/20重叠 ● VPC路由表中自定义路由的目 的地址不能与 198.19.128.0/20重叠

	SubnetId *string `json:"subnet_id,omitempty"`
	// 终端节点服务的ID。 可以通过查询终端节点服务概 要获取要连接的终端节点服务 ID。

	EndpointServiceId string `json:"endpoint_service_id"`
	// 终端节点所在的VPC的ID。 详细内容请参考《虚拟私有云 API参考》中的“查询VPC”， 详见响应消息中的“id”字 段。

	VpcId string `json:"vpc_id"`
	// 是否创建域名。 ● true：创建域名 ● false：不创建域名 默认值为false。 说明 当创建连接gateway类型终端节点服务的终端节点时，“enable_dns”设置为true或者false，均不创建域名。

	EnableDns *bool `json:"enable_dns,omitempty"`
	// 标签列表，没有标签默认为空数组。

	Tags *[]TagList `json:"tags,omitempty"`
	// 路由表ID列表。详细内容请参考《虚拟私有云 API参考》中的“查询VPC路由”，详见响应消息中的“id”字段。 创建连接gateway类型终端节点服务的终节点时，此参数必选。 说明 不设置此参数时，选择默认路由表。

	Routetables *[]string `json:"routetables,omitempty"`
	// 访问所连接的终端节点服务的IP。 创建终端节点时，可以指定访问所连接的终端节点服务的IP，目前只支持IPv4类型 。 创建连接Interface类型终端节点服务的终端节点时，此参数必选。

	PortIp *string `json:"port_ip,omitempty"`
	// 添加用于控制访问终端节点的白名单。 创建终端节点时，支持访问控制，使用此参数可以添加IPv4或CIDR，默认空列表。 仅当创建连接Interface类型终端节点服务的终端节点时，支持设置此参数。

	Whitelist *[]string `json:"whitelist,omitempty"`
	// 是否开启网络ACL隔离。

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`
}

func (o CreateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequestBody", string(data)}, " ")
}
