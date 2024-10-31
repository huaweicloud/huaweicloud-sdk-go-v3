package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFirewallRequest Request Object
type DeleteFirewallRequest struct {

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	ResourceId string `json:"resource_id"`
}

func (o DeleteFirewallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFirewallRequest struct{}"
	}

	return strings.Join([]string{"DeleteFirewallRequest", string(data)}, " ")
}
