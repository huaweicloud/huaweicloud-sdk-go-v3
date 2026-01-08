package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainIPsResponse Response Object
type ListDomainIPsResponse struct {

	// 负载均衡器dns ip信息列表。
	Ips            *[]DnsIpResponse `json:"ips,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListDomainIPsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainIPsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainIPsResponse", string(data)}, " ")
}
