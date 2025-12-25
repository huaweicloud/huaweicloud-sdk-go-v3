package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CollectorNode struct {

	// 数值
	ChannelInstanceReferCount *int64 `json:"channel_instance_refer_count,omitempty"`

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// 自定义标签
	CustomLabel *string `json:"custom_label,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 设备类型
	DeviceType *string `json:"device_type,omitempty"`

	// IP地址
	IpAddress *string `json:"ip_address,omitempty"`

	Monitor *Monitor `json:"monitor,omitempty"`

	NodeExpansion *IsapNodeExpansion `json:"node_expansion,omitempty"`

	// UUID
	NodeId *string `json:"node_id,omitempty"`

	// 所属租户名称
	NodeName *string `json:"node_name,omitempty"`

	// 操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// IP地址
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 规格
	Specification *string `json:"specification,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// Vpc 端点地址
	VpcEndpointAddress *string `json:"vpc_endpoint_address,omitempty"`

	// Vpc 端点ID
	VpcEndpointId *string `json:"vpc_endpoint_id,omitempty"`

	// UUID
	VpcId *string `json:"vpc_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o CollectorNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectorNode struct{}"
	}

	return strings.Join([]string{"CollectorNode", string(data)}, " ")
}
