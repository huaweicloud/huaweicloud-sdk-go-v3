package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDnsIpResponseBody 负载均衡器dns ip信息列表。
type ListDnsIpResponseBody struct {

	// 负载均衡器dns ip信息列表。
	Ips *[]DnsIpResponse `json:"ips,omitempty"`
}

func (o ListDnsIpResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDnsIpResponseBody struct{}"
	}

	return strings.Join([]string{"ListDnsIpResponseBody", string(data)}, " ")
}
