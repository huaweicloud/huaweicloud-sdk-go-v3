package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeerAddress struct {

	// 域名地址
	DomainPort *string `json:"domain_port,omitempty"`

	// IP地址
	IpPort *string `json:"ip_port,omitempty"`
}

func (o PeerAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerAddress struct{}"
	}

	return strings.Join([]string{"PeerAddress", string(data)}, " ")
}
