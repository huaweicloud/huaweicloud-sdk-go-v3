package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性公网IP字典对象
type PublicIp struct {

	// 弹性公网IP唯一标识。
	Id *string `json:"id,omitempty" xml:"id"`

	// 弹性公网IP的状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 端口的ID。
	PortId *string `json:"port_id,omitempty" xml:"port_id"`

	// 弹性公网IP的地址。
	PublicIpAddress *string `json:"public_ip_address,omitempty" xml:"public_ip_address"`

	// 绑定弹性公网IP的私有IP地址。
	PrivateIpAddress *string `json:"private_ip_address,omitempty" xml:"private_ip_address"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 带宽的ID。
	BandwidthId *string `json:"bandwidth_id,omitempty" xml:"bandwidth_id"`

	// 带宽的名称。
	BandwidthName *string `json:"bandwidth_name,omitempty" xml:"bandwidth_name"`

	// 带宽的类型。
	BandwidthShareType *string `json:"bandwidth_share_type,omitempty" xml:"bandwidth_share_type"`

	// 带宽的大小。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty" xml:"bandwidth_size"`

	// IP版本的信息。
	IpVersion *int32 `json:"ip_version,omitempty" xml:"ip_version"`

	// 子网所属的站点ID。
	SiteId *string `json:"site_id,omitempty" xml:"site_id"`

	// 子网所属的站点信息。
	SiteInfo *string `json:"site_info,omitempty" xml:"site_info"`

	Operator *Operator `json:"operator,omitempty" xml:"operator"`

	// 弹性公网IP的类型。
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
