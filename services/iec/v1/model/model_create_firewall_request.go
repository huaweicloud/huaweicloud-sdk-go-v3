package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateFirewallRequest struct {
	Body *CreateFirewallRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateFirewallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallRequest struct{}"
	}

	return strings.Join([]string{"CreateFirewallRequest", string(data)}, " ")
}
