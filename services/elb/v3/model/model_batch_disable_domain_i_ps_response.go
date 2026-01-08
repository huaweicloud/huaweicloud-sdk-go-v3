package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableDomainIPsResponse Response Object
type BatchDisableDomainIPsResponse struct {

	// **参数解释**：返回的负载均衡器的dns ip信息。  **约束限制**：如果负载均衡器的公网域名和私网域名开关都没有打开，则列表为空。  **取值范围**：不涉及  **默认取值**：不涉及
	Ips            *[]ListDnsIpResponseBody `json:"ips,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchDisableDomainIPsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableDomainIPsResponse struct{}"
	}

	return strings.Join([]string{"BatchDisableDomainIPsResponse", string(data)}, " ")
}
