package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkAvailableIpResponse Response Object
type ShowNetworkAvailableIpResponse struct {

	// **参数解释**：子网的名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：子网的ID。 **取值范围**：不涉及。
	NetworkId *string `json:"networkId,omitempty"`

	// **参数解释**：子网可用的网络IP数量。
	SubnetIpAvailability *[]SubnetIpAvailability `json:"subnetIpAvailability,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ShowNetworkAvailableIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkAvailableIpResponse struct{}"
	}

	return strings.Join([]string{"ShowNetworkAvailableIpResponse", string(data)}, " ")
}
