package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NeutronDeleteFloatingIpRequest struct {
	// floatingip的ID

	FloatingipId string `json:"floatingip_id"`
}

func (o NeutronDeleteFloatingIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFloatingIpRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFloatingIpRequest", string(data)}, " ")
}
