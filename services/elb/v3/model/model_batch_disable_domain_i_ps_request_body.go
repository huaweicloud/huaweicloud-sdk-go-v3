package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableDomainIPsRequestBody 批量从负载均衡器域名解析中移除IP地址的请求体。
type BatchDisableDomainIPsRequestBody struct {

	// **参数解释**：需要从负载均衡器域名解析中移除的IP列表。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Ips []DnsIp `json:"ips"`
}

func (o BatchDisableDomainIPsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableDomainIPsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDisableDomainIPsRequestBody", string(data)}, " ")
}
