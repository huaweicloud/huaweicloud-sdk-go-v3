package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdatePortRequest Request Object
type NeutronUpdatePortRequest struct {

	// 端口ID
	PortId string `json:"port_id"`

	Body *NeutronUpdatePortRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdatePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdatePortRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdatePortRequest", string(data)}, " ")
}
