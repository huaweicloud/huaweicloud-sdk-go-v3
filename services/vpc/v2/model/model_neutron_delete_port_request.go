package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeletePortRequest Request Object
type NeutronDeletePortRequest struct {

	// 端口ID
	PortId string `json:"port_id"`
}

func (o NeutronDeletePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeletePortRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeletePortRequest", string(data)}, " ")
}
