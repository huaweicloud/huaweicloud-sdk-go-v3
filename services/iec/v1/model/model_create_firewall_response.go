package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateFirewallResponse struct {
	Firewall       *Firewall `json:"firewall,omitempty" xml:"firewall"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallResponse struct{}"
	}

	return strings.Join([]string{"CreateFirewallResponse", string(data)}, " ")
}
