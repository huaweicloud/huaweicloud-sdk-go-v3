package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateFirewallResponse struct {
	Firewall       *UpdateFirewallResp `json:"firewall,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallResponse struct{}"
	}

	return strings.Join([]string{"UpdateFirewallResponse", string(data)}, " ")
}
