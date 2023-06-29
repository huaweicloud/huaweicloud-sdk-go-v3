package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateFloatingIpRequest Request Object
type NeutronUpdateFloatingIpRequest struct {

	// floatingip的ID
	FloatingipId string `json:"floatingip_id"`

	Body *NeutronUpdateFloatingIpRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateFloatingIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFloatingIpRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFloatingIpRequest", string(data)}, " ")
}
