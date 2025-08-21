package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeEastWestFirewallStatusRequest Request Object
type ChangeEastWestFirewallStatusRequest struct {
	Body *ChangeProtectStatusRequestBody `json:"body,omitempty"`
}

func (o ChangeEastWestFirewallStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEastWestFirewallStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeEastWestFirewallStatusRequest", string(data)}, " ")
}
