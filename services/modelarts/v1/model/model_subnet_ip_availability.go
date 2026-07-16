package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetIpAvailability 子网可用的IP信息。
type SubnetIpAvailability struct {

	// **参数解释**：子网的cidr。 **取值范围**：不涉及。
	Cidr string `json:"cidr"`

	// **参数解释**：网络版本。 **取值范围**：可选值如下： - 4：代表ipV4
	IpVersion int32 `json:"ipVersion"`

	// **参数解释**：已使用的IP数量。 **取值范围**：不涉及。
	UsedIps int32 `json:"usedIps"`

	// **参数解释**：子网中总的IP数量。 **取值范围**：不涉及。
	TotalIps *int32 `json:"totalIps,omitempty"`
}

func (o SubnetIpAvailability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetIpAvailability struct{}"
	}

	return strings.Join([]string{"SubnetIpAvailability", string(data)}, " ")
}
