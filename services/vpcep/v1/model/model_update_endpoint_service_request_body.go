package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改终端节点服务接口请求结构体
type UpdateEndpointServiceRequestBody struct {
	// 是否需要审批。 ● false：不需审批，创建的终端节点连接直接为accepted状态。 ● true：需审批，创建的终端节点连接需要终端节点服务所属用户审核后方可使用。 默认为true，需要审批。

	ApprovalEnabled *bool `json:"approval_enabled,omitempty"`
	// 终端节点服务的名称，长度不大于16，允许传入大小写字母、数字、下划线、中划线。

	ServiceName *string `json:"service_name,omitempty"`
	// 服务开放的端口映射列表，详细内容请参见表4-22。 同一个终端节点服务下，不允许重复的端口映射。若多个终端节点服务共用一个port_id，则终端节点之间服务的所有端口映射的server_port和protocol的组合不能重复，单次最多添加200个。

	Ports *[]PortList `json:"ports,omitempty"`
	// 标识终端节点服务后端资源的ID，格式为通用唯一识别码 （Universally UniqueIdentifier，下文简称UUID）。取值为： ● LB类型：增强型负载均衡器内网IP对应的端口ID。详细内容请参考《弹性负载均衡API参考》中的“查询负载均 衡详情”，详见响应消息中的“vip_port_id”字段。 ● VM类型：弹性云服务器IP地址对应的网卡ID。详细内容请参考《弹性云服务器API参考》中的“查询云服务器网 卡信息”，详见响应消息中的“port_id”字段。 ● VIP类型：虚拟资源所在物理服务器对应的网卡ID。 说明 当后端资源为“LB类型”时，仅支持修改为同类型后端资源的“vip_port_id”。 例如，共享型负载均衡仅支持更换为共享型负载均衡，不支持更换为独享型负载均衡。

	PortId *string `json:"port_id,omitempty"`
	// 虚拟IP的网卡ID。

	VipPortId *string `json:"vip_port_id,omitempty"`
}

func (o UpdateEndpointServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceRequestBody", string(data)}, " ")
}
