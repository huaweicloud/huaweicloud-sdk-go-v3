package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableDomainIPsResponse Response Object
type BatchEnableDomainIPsResponse struct {

	// **参数解释**：负载均衡器域名解析的IP地址列表。  **约束限制**：如果负载均衡器的公网域名和私网域名域名解析开关都没有打开，则为空列表。  **取值范围**：不涉及  **默认取值**：不涉及
	Ips *[]DnsIpResponse `json:"ips,omitempty"`

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchEnableDomainIPsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableDomainIPsResponse struct{}"
	}

	return strings.Join([]string{"BatchEnableDomainIPsResponse", string(data)}, " ")
}
