package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListPortsRequest Request Object
type NeutronListPortsRequest struct {

	// 每页返回的个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// 按照端口的ID过滤查询
	Id *string `json:"id,omitempty"`

	// 按照端口的名称过滤查询
	Name *string `json:"name,omitempty"`

	// 按照端口的管理状态过滤查询，取值范围：true or false
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 按照端口所属的网络ID过滤查询
	NetworkId *string `json:"network_id,omitempty"`

	// 按照端口的mac地址过滤查询
	MacAddress *string `json:"mac_address,omitempty"`

	// 按照端口的设备ID过滤查询
	DeviceId *string `json:"device_id,omitempty"`

	// 按照端口的设备所属过滤查询
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 按照端口状态过滤查询，取值范围：ACTIVE、BUILD、DOWN
	Status *string `json:"status,omitempty"`

	// 按照安全组ID列表过滤查询
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// 按照端口的IP地址过滤查询，fixed_ips=ip_address或者fixed_ips=subnet_id过滤查询
	FixedIps *[]string `json:"fixed_ips,omitempty"`

	// 按照端口所属的项目ID过滤查询
	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NeutronListPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListPortsRequest struct{}"
	}

	return strings.Join([]string{"NeutronListPortsRequest", string(data)}, " ")
}
