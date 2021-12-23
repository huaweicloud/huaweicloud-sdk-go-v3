package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronShowFloatingIpResponse struct {
	Floatingip     *FloatingIpResp `json:"floatingip,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o NeutronShowFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowFloatingIpResponse", string(data)}, " ")
}
