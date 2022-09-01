package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type SubnetIpAvailability struct {

	// 子网中已经使用的IP数目（不包含系统预留地址）
	UsedIps int32 `json:"used_ips" xml:"used_ips"`

	// 子网ID
	SubnetId string `json:"subnet_id" xml:"subnet_id"`

	// 子网名称
	SubnetName string `json:"subnet_name" xml:"subnet_name"`

	// 子网的IP版本，取值为4或者6
	IpVersion int32 `json:"ip_version" xml:"ip_version"`

	// 子网的CIDR
	Cidr string `json:"cidr" xml:"cidr"`

	// 子网中IP总数（不包含系统预留地址）
	TotalIps int32 `json:"total_ips" xml:"total_ips"`
}

func (o SubnetIpAvailability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetIpAvailability struct{}"
	}

	return strings.Join([]string{"SubnetIpAvailability", string(data)}, " ")
}
