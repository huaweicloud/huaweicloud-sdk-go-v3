package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteFirewallResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFirewallResponse struct{}"
	}

	return strings.Join([]string{"DeleteFirewallResponse", string(data)}, " ")
}
