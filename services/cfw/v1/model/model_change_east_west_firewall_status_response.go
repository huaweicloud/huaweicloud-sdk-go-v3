package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeEastWestFirewallStatusResponse Response Object
type ChangeEastWestFirewallStatusResponse struct {
	Data *SuccessRspData `json:"data,omitempty"`

	// trace id
	TraceId        *string `json:"trace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeEastWestFirewallStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEastWestFirewallStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeEastWestFirewallStatusResponse", string(data)}, " ")
}
