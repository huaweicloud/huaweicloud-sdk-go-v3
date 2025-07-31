package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomerIpsRequest Request Object
type UpdateCustomerIpsRequest struct {

	// **参数解释**： cfw侧自定义IPS规则id **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	IpsCfwId string `json:"ips_cfw_id"`

	Body *CustomerIpsSaveDto `json:"body,omitempty"`
}

func (o UpdateCustomerIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomerIpsRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomerIpsRequest", string(data)}, " ")
}
