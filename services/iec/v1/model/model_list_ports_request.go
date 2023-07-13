package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsRequest Request Object
type ListPortsRequest struct {

	// 查询返回端口列表数量。取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 子网的neutron的network的ID。
	NetworkId *string `json:"network_id,omitempty"`

	// 按照端口ID过滤查询
	Id *string `json:"id,omitempty"`

	// 按照name过滤查询  取值范围：最大长度不超过255
	Name *string `json:"name,omitempty"`

	// 按照admin_state_up进行过滤  约束：只支持true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 根据绑定的IP查询端口。按照fixed_ips=ip_address或者fixed_ips=subnet_id过滤查询，示例：fixed_ips=ip_address=xxx&fixed_ips=subnet_id=xxxx
	FixedIps *[]string `json:"fixed_ips,omitempty"`

	// 根据网卡的mac地址查询端口。
	MacAddress *string `json:"mac_address,omitempty"`

	// 根据设备ID查询端口。
	DeviceId *string `json:"device_id,omitempty"`

	// 根据设备主查询端口。
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 按照status过滤查询  取值范围：ACTIVE、BUILD、DOWN
	Status *string `json:"status,omitempty"`

	// 根据安全组信息ID查询端口。
	SecurityGroups *string `json:"security_groups,omitempty"`
}

func (o ListPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsRequest struct{}"
	}

	return strings.Join([]string{"ListPortsRequest", string(data)}, " ")
}
