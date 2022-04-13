package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronDeleteFloatingIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFloatingIpResponse", string(data)}, " ")
}
