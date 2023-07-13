package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Network struct {

	// 方向，取值范围：IN、OUT。
	Direction *string `json:"direction,omitempty"`

	// 协议。
	Protocol *string `json:"protocol,omitempty"`

	// 源IP地址。
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口，0–65535。
	SrcPort *int32 `json:"src_port,omitempty"`

	// 源域名，最大128个字符。
	SrcDomain *string `json:"src_domain,omitempty"`

	SrcGeo *Geo `json:"src_geo,omitempty"`

	// 目标IP地址。
	DestIp *string `json:"dest_ip,omitempty"`

	// 目标端口，0–65535。
	DestPort *int32 `json:"dest_port,omitempty"`

	// 目标域名，最大128个字符。
	DestDomain *string `json:"dest_domain,omitempty"`

	DestGeo *Geo `json:"dest_geo,omitempty"`
}

func (o Network) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Network struct{}"
	}

	return strings.Join([]string{"Network", string(data)}, " ")
}
