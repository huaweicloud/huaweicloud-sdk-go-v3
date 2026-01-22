package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFirewallRequest Request Object
type DeleteFirewallRequest struct {

	// **参数解释**： 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ResourceId string `json:"resource_id"`
}

func (o DeleteFirewallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFirewallRequest struct{}"
	}

	return strings.Join([]string{"DeleteFirewallRequest", string(data)}, " ")
}
