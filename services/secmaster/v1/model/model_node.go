package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Node struct {

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	Description *IsapErrorRsp `json:"description,omitempty"`

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

	// 区域
	Region *string `json:"region,omitempty"`

	// 规格
	Specification *string `json:"specification,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// Vpc 端点地址
	VpcEndpointAddress *string `json:"vpc_endpoint_address,omitempty"`

	// Vpc 端点ID
	VpcEndpointId *string `json:"vpc_endpoint_id,omitempty"`

	// UUID
	VpcId *string `json:"vpc_id,omitempty"`

	// IP地址
	VpcepServiceIp *string `json:"vpcep_service_ip,omitempty"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}
