package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建floatingip对象
type NeutronCreateFloatingIpRequestBody struct {
	Floatingip *CreateFloatingIpOption `json:"floatingip"`
}

func (o NeutronCreateFloatingIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFloatingIpRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFloatingIpRequestBody", string(data)}, " ")
}
