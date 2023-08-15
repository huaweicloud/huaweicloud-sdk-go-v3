package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAlertRspNetworkList struct {

	// 方向，取值范围：IN | OUT
	Direction *interface{} `json:"direction,omitempty"`

	// 协议，参考：IANA registered name
	Protocol *string `json:"protocol,omitempty"`

	// 源IP地址
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口，0–65535
	SrcPort *string `json:"src_port,omitempty"`

	// 源域名，最大128个字符
	SrcDomain *string `json:"src_domain,omitempty"`

	// 目的IP地址
	DestIp *string `json:"dest_ip,omitempty"`

	// 目的端口，0–65535
	DestPort *string `json:"dest_port,omitempty"`

	// 目的域名，最大128个字符
	DestDomain *string `json:"dest_domain,omitempty"`

	// 源IP的地理位置信息
	SrcGeo *interface{} `json:"src_geo,omitempty"`

	// 目的IP的地理位置信息
	DestGeo *interface{} `json:"dest_geo,omitempty"`
}

func (o ShowAlertRspNetworkList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRspNetworkList struct{}"
	}

	return strings.Join([]string{"ShowAlertRspNetworkList", string(data)}, " ")
}
