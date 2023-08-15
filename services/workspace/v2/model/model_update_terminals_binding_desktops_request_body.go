package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTerminalsBindingDesktopsRequestBody struct {

	// 策略id
	Id string `json:"id"`

	// 终端MAC地址
	Mac string `json:"mac"`

	// 虚拟机名称
	DesktopName string `json:"desktop_name"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateTerminalsBindingDesktopsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTerminalsBindingDesktopsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTerminalsBindingDesktopsRequestBody", string(data)}, " ")
}
