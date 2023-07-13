package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowPortRequest Request Object
type NeutronShowPortRequest struct {

	// 端口ID
	PortId string `json:"port_id"`
}

func (o NeutronShowPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowPortRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowPortRequest", string(data)}, " ")
}
