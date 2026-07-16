package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkAvailableIpRequest Request Object
type ShowNetworkAvailableIpRequest struct {

	// **参数解释**：网络ID。取值自网络详情的metadata.name字段。 **约束限制**：1 - 64字符，字母、数字和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NetworkName string `json:"network_name"`

	// **参数解释**：子网ID。取值自网络详情status.subnets字段中的networkId字段。 **约束限制**：1 - 64字符，字母、数字和中划线。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NetworkId string `json:"network_id"`
}

func (o ShowNetworkAvailableIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkAvailableIpRequest struct{}"
	}

	return strings.Join([]string{"ShowNetworkAvailableIpRequest", string(data)}, " ")
}
