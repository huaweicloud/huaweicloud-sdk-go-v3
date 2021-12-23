package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronUpdateFloatingIpResponse struct {
	Floatingip     *PostAndPutFloatingIpResp `json:"floatingip,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o NeutronUpdateFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFloatingIpResponse", string(data)}, " ")
}
