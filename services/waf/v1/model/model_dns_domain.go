package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnsDomain struct {

	// dns-id
	Id *string `json:"id,omitempty"`

	// 域名
	Domain *string `json:"domain,omitempty"`

	// dns服务信息
	Servers *[]Server `json:"servers,omitempty"`

	// 防护端口
	ProtectPort *string `json:"protect_port,omitempty"`
}

func (o DnsDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsDomain struct{}"
	}

	return strings.Join([]string{"DnsDomain", string(data)}, " ")
}
