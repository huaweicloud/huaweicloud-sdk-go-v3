package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableDomainIPsRequestBody 批量向负载均衡器域名解析中加入IP地址的请求体。
type BatchEnableDomainIPsRequestBody struct {

	// **参数解释**：需要加入到负载均衡器域名解析中的IP列表。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Ips []DnsIp `json:"ips"`
}

func (o BatchEnableDomainIPsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableDomainIPsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchEnableDomainIPsRequestBody", string(data)}, " ")
}
