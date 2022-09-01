package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPortsRequest struct {

	// 查询返回端口列表数量。取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 子网的neutron的network的ID。
	NetworkId *string `json:"network_id,omitempty" xml:"network_id"`

	// 按照端口ID过滤查询
	Id *string `json:"id,omitempty" xml:"id"`

	// 按照name过滤查询  取值范围：最大长度不超过255
	Name *string `json:"name,omitempty" xml:"name"`

	// 按照admin_state_up进行过滤  约束：只支持true
	AdminStateUp *bool `json:"admin_state_up,omitempty" xml:"admin_state_up"`

	// 根据绑定的IP查询端口。按照fixed_ips=ip_address或者fixed_ips=subnet_id过滤查询，示例：fixed_ips=ip_address=xxx&fixed_ips=subnet_id=xxxx
	FixedIps *[]string `json:"fixed_ips,omitempty" xml:"fixed_ips"`

	// 根据网卡的mac地址查询端口。
	MacAddress *string `json:"mac_address,omitempty" xml:"mac_address"`

	// 根据设备ID查询端口。
	DeviceId *string `json:"device_id,omitempty" xml:"device_id"`

	// 根据设备主查询端口。
	DeviceOwner *string `json:"device_owner,omitempty" xml:"device_owner"`

	// 按照status过滤查询  取值范围：ACTIVE、BUILD、DOWN
	Status *string `json:"status,omitempty" xml:"status"`

	// 根据安全组信息ID查询端口。
	SecurityGroups *string `json:"security_groups,omitempty" xml:"security_groups"`
}

func (o ListPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsRequest struct{}"
	}

	return strings.Join([]string{"ListPortsRequest", string(data)}, " ")
}
