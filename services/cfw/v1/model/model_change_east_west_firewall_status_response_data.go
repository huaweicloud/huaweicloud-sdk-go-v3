package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeEastWestFirewallStatusResponseData struct {

	// **参数解释**： VPC边界防护对象id **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o ChangeEastWestFirewallStatusResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEastWestFirewallStatusResponseData struct{}"
	}

	return strings.Join([]string{"ChangeEastWestFirewallStatusResponseData", string(data)}, " ")
}
