package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainIPsResponse Response Object
type ListDomainIPsResponse struct {

	// 负载均衡器IP地址的域名解析配置列表。
	Ips *[]DnsIpResponse `json:"ips,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainIPsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainIPsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainIPsResponse", string(data)}, " ")
}
