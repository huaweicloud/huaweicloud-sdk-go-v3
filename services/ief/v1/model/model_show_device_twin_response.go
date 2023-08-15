package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceTwinResponse Response Object
type ShowDeviceTwinResponse struct {
	PropertyVisitors *ValueInPropertyVisitors `json:"property_visitors,omitempty"`

	Twin *ValueInTwinResponse `json:"twin,omitempty"`

	// 访问协议，有如下选项： - userdefine：自定义协议 - modbus：modbus协议 - opc-ua：opc-ua协议
	AccessProtocol *string `json:"access_protocol,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDeviceTwinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceTwinResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceTwinResponse", string(data)}, " ")
}
