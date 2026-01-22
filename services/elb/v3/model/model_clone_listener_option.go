package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneListenerOption 监听器复制到实例相关配置参数。
type CloneListenerOption struct {

	// **参数解释**：新监听器的名称。 **约束限制**：不涉及 **取值范围**：[0-255]个字符，传入空字符串则取默认值。 **默认取值**：原监听器名称+ “-copy”。
	Name *string `json:"name,omitempty"`

	// **参数解释**：新监听器所在的负载均衡器ID（UUID）。 **约束限制**： - 不支持复制到网关型负载均衡器。 - 目的负载均衡器类型需要支持源监听器协议。源监听器协议为HTTP、HTTPS，目的负载均衡器需要为应用型负载均衡器；源监听器协议为TCP、UDP、TLS，目的负载均衡器需要为网络型负载均衡器。 - 只支持共享型负载均衡器复制到共享型负载均衡器，独享型负载均衡器复制到独享型负载均衡器。 - 源监听器的负载均衡器启用IP类型后端转发，目的负载均衡器也需要启动IP类型后端转发功能。 - 源监听器协议为TLS，目的负载均衡器需要支持创建TLS协议监听器。 **取值范围**：标准的UUID格式，长度为36个字符。 **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	// **参数解释**：新监听器的监听端口。 **约束限制**： - 不能与目的负载均衡器下已有监听器监听端口重复。 - 0表示开启监听端口范围的能力，此时port_ranges为必填字段。 - 共享型负载均衡器上的HTTP和TERMINATED_HTTPS监听器不支持设置21端口。 **取值范围**：0-65535 **默认取值**：不涉及
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// **参数解释**：端口监听范围（闭区间)，最多指定10个端口组，每个组范围不可有重叠部分。 **约束限制**： - 仅当protocol_port为0或未传入protocol_port时可以传入该字段。 - 仅TCP, UDP，TLS监听器支持该字段。 - 不能与目的负载均衡器下已有端口冲突
	PortRanges *[]PortRange `json:"port_ranges,omitempty"`

	// **参数解释**：新监听器是否复用或复制源监听器的后端服务器组和后端服务器的标识。 - 复用：目的监听器将会直接使用源监听器的后端服务器组。 - 复制：系统将会根据原有配置创建新的后端服务器组，然后将其关联到目的监听器使用。 **约束限制**： - 配置为true时，需要开启后端服务器组多挂实例功能。 - 只在独享型场景复制、同VPC场景可配。 **取值范围**： - true：复用源监听器的后端服务器组。 - false：复制源监听器的后端服务器组。 **默认取值**：false
	ReusePool *bool `json:"reuse_pool,omitempty"`

	// **参数解释**：源监听器下后端服务器子网信息和新监听器下后端服务器子网信息一一对应关系。 **约束限制**： - 将监听器复制到不同VPC下的负载均衡器时，该字段必填。复制到同一个VPC下的负载均衡器时不填。 - 若源监听器所在负载均衡器已开启ip_target_enable（该功能默认不开启），则不允许跨VPC复制，即该字段不允许填。 - 每一组subnet_cidr_id都需要是新监听器下后端服务器的VPC子网ID，每一组dst_subnet_cidr_id都需要为源监听器下后端服务器的的VPC子网ID，不允许少填多填、或重复对应关系。 - 每一组的subnet_cidr_id和dst_subnet_cidr_id的两个子网必须存在且网段相同。
	SubnetMappingList *[]SubnetMappingList `json:"subnet_mapping_list,omitempty"`
}

func (o CloneListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneListenerOption struct{}"
	}

	return strings.Join([]string{"CloneListenerOption", string(data)}, " ")
}
