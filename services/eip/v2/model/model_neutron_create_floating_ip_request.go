package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateFloatingIpRequest Request Object
type NeutronCreateFloatingIpRequest struct {
	Body *NeutronCreateFloatingIpRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateFloatingIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFloatingIpRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateFloatingIpRequest", string(data)}, " ")
}
