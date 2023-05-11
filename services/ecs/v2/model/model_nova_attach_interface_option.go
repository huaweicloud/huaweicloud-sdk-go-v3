package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NovaAttachInterfaceOption struct {

	// 私有IP。  有port_id时，该参数不起作用，并且有该参数时，必须有net_id。  只有列表中第一个元素有效。传入两个及以上元素会报错。
	FixedIps *[]NovaAttachInterfaceFixedIp `json:"fixed_ips,omitempty"`

	//   Network ID。  有port_id时，该参数不起作用。
	NetId *string `json:"net_id,omitempty"`

	//   Port ID。  port_id和net_id不能同时传入。
	PortId *string `json:"port_id,omitempty"`
}

func (o NovaAttachInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceOption struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceOption", string(data)}, " ")
}
