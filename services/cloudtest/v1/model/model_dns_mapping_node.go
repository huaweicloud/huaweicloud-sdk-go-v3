package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DnsMappingNode DNS配置
type DnsMappingNode struct {

	// 域名信息
	DomainName *string `json:"domain_name,omitempty"`

	// 域名对应的IP
	Ips *[]string `json:"ips,omitempty"`
}

func (o DnsMappingNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsMappingNode struct{}"
	}

	return strings.Join([]string{"DnsMappingNode", string(data)}, " ")
}
