package model

import (
	"encoding/json"

	"strings"
)

type CdnIps struct {
	// 需查询的IP地址。

	Ip *string `json:"ip,omitempty"`
	// 是否是华为云CDN节点。

	Belongs *string `json:"belongs,omitempty"`
	// IP归属地。

	Region *string `json:"region,omitempty"`
	// 运营商。

	Isp *string `json:"isp,omitempty"`
	// 平台。

	Platform *string `json:"platform,omitempty"`
}

func (o CdnIps) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CdnIps struct{}"
	}

	return strings.Join([]string{"CdnIps", string(data)}, " ")
}
