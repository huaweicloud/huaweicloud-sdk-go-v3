package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateFloatingIpResponse Response Object
type NeutronCreateFloatingIpResponse struct {
	Floatingip     *PostAndPutFloatingIpResp `json:"floatingip,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o NeutronCreateFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateFloatingIpResponse", string(data)}, " ")
}
