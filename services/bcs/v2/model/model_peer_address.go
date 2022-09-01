package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeerAddress struct {

	// 域名地址
	DomainPort *string `json:"domain_port,omitempty" xml:"domain_port"`

	// IP地址
	IpPort *string `json:"ip_port,omitempty" xml:"ip_port"`
}

func (o PeerAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerAddress struct{}"
	}

	return strings.Join([]string{"PeerAddress", string(data)}, " ")
}
