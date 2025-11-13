package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyFlavor struct {

	// **参数解释**：  规格ID。  **取值范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  规格码。  **取值范围**：  不涉及。
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**：  cpu核数。  **取值范围**：  不涉及。
	Vcpus *string `json:"vcpus,omitempty"`

	// **参数解释**：  内存。  **取值范围**：  不涉及。
	Ram *string `json:"ram,omitempty"`

	// **参数解释**：  数据库类型。  **取值范围**：  不涉及。
	DbType *string `json:"db_type,omitempty"`

	// **参数解释**：  az状态。
	AzStatus *interface{} `json:"az_status,omitempty"`

	// **参数解释**：  是否支持ipv6。  **取值范围**：  - true: 支持ipv6。 - false: 不支持ipv6。
	SupportedIpv6 *bool `json:"supported_ipv6,omitempty"`
}

func (o ProxyFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyFlavor struct{}"
	}

	return strings.Join([]string{"ProxyFlavor", string(data)}, " ")
}
