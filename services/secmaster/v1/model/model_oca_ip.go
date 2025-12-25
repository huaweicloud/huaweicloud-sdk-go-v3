package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIp 线下机房ip
type OcaIp struct {

	// 资产值
	Value string `json:"value"`

	// 资产类型：  ipv4、ipv6
	Version string `json:"version"`

	Network *OcaIpNetwork `json:"network"`

	// 资产备注
	Remark *string `json:"remark,omitempty"`

	// 资产名称，默认为资产值
	Name *string `json:"name,omitempty"`

	// 相对值，如资产为ipv4，则为对应的ipv6
	RelativeValue *string `json:"relative_value,omitempty"`

	// 机房
	ServerRoom string `json:"server_room"`

	// 机柜
	ServerRack string `json:"server_rack"`

	DataCenter *OcaIpDataCenter `json:"data_center"`

	// mac地址
	MacAddr *string `json:"mac_addr,omitempty"`

	// 重要等级0 ：不重要 1：重要
	Important *string `json:"important,omitempty"`

	ExtendProperties *OcaIpExtendProperties `json:"extend_properties,omitempty"`
}

func (o OcaIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIp struct{}"
	}

	return strings.Join([]string{"OcaIp", string(data)}, " ")
}
