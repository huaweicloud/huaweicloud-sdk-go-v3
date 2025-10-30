package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentAddressResponse Response Object
type ShowAgentAddressResponse struct {

	// **参数解释**： pod_lb地址 **取值范围**： 字符长度1-256位
	AnpAddress *string `json:"anp_address,omitempty"`

	// **参数解释**： region_id标识 **取值范围**： 字符长度1-128位
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释**： 公网接入agent地址 **取值范围**： 字符长度1-256位
	AgentAddress   *string `json:"agent_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAgentAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentAddressResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentAddressResponse", string(data)}, " ")
}
