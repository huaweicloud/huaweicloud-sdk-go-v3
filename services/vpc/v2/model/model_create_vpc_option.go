package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcOption vpc对象
type CreateVpcOption struct {

	// 功能说明：虚拟私有云下可用子网的范围 取值范围： - 10.0.0.0/8 ~ 10.255.255.240/28 - 172.16.0.0/12 ~ 172.31.255.240/28 - 192.168.0.0/16 ~ 192.168.255.240/28 约束：必须是ipv4 cidr格式，例如:192.168.0.0/16
	Cidr *string `json:"cidr,omitempty"`

	// 功能说明：虚拟私有云名称 取值范围：0-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点） 约束：如果名称不为空，则同一个租户下的名称不能重复
	Name *string `json:"name,omitempty"`

	// 功能说明：虚拟私有云的描述 取值范围：0-255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 功能说明：企业项目ID。创建虚拟私有云时，给虚拟私有云绑定企业项目ID。 取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 默认值：\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 功能说明：VPC资源标签。创建VPC时，给VPC添加资源标签。 取值范围：最大10个标签, key：标签名称; value：标签值。 格式：[key*value]，每一个标签的key和value之间用*连接
	Tags *[]string `json:"tags,omitempty"`

	// 功能说明：默认情况下，VPC中的资源可以通过内网访问服务终结点。开启该项后，VPC将无法通过内网访问服务终结点，请谨慎操作。 无法访问以下云服务：容器镜像服务SWR、云日志服务LTS、企业主机安全HSS、应用运维管理AOM、应用性能管理APM、对象存储服务OBS、API网关APIG。 取值范围： off：代表禁用。 on：代表开启。
	BlockServiceEndpointStates *string `json:"block_service_endpoint_states,omitempty"`

	// 功能说明：是否开启VPC内所有子网的IPv4地址使用量指标监控。 取值范围： true：开启 false：不开启
	EnableNetworkAddressUsageMetrics *bool `json:"enable_network_address_usage_metrics,omitempty"`
}

func (o CreateVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcOption struct{}"
	}

	return strings.Join([]string{"CreateVpcOption", string(data)}, " ")
}
