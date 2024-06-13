package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DnsMapping DNS配置结果集
type DnsMapping struct {
	Dns *[]DnsMappingNode `json:"dns,omitempty"`
}

func (o DnsMapping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsMapping struct{}"
	}

	return strings.Join([]string{"DnsMapping", string(data)}, " ")
}
