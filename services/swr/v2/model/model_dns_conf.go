package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnsConf struct {

	// 主机映射，键为域名，值为对应的 IP 地址
	Hosts map[string]string `json:"hosts,omitempty"`
}

func (o DnsConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsConf struct{}"
	}

	return strings.Join([]string{"DnsConf", string(data)}, " ")
}
