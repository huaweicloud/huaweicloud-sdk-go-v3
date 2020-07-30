/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type RejectVpcPeeringResponse struct {
	// 对等连接ID
	Id string `json:"id,omitempty"`
	// 功能说明：对等连接名称 取值范围：支持1~64个字符
	Name string `json:"name,omitempty"`
	// 功能说明：对等连接状态 取值范围： - PENDING_ACCEPTANCE：等待接受 - REJECTED：已拒绝。 - EXPIRED：已过期。 - DELETED：已删除。 - ACTIVE：活动的。
	Status string `json:"status,omitempty"`
	RequestVpcInfo *VpcInfo `json:"request_vpc_info,omitempty"`
	AcceptVpcInfo *VpcInfo `json:"accept_vpc_info,omitempty"`
	// 功能说明：资源创建UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	CreatedAt string `json:"created_at,omitempty"`
	// 功能说明：资源更新UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	UpdatedAt string `json:"updated_at,omitempty"`
	// 对等连接描述
	Description string `json:"description,omitempty"`
}
