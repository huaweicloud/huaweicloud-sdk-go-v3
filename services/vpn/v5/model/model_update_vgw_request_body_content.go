package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVgwRequestBodyContent struct {

	// 网关名称
	Name *string `json:"name,omitempty"`

	// 本端子网
	LocalSubnets *[]string `json:"local_subnets,omitempty"`

	// 主eip的ID。用于给VPN网关绑定新的主EIP，需要先解绑当前的主EIP
	MasterEipId *string `json:"master_eip_id,omitempty"`

	// 备eip的ID。用于给VPN网关绑定新的备EIP，需要先解绑当前的备EIP
	SlaveEipId *string `json:"slave_eip_id,omitempty"`

	// 主eip的ID。用于给VPN网关绑定新的主EIP，需要先解绑当前的主EIP
	EipId1 *string `json:"eip_id_1,omitempty"`

	// 备eip的ID。用于给VPN网关绑定新的备EIP，需要先解绑当前的备EIP
	EipId2 *string `json:"eip_id_2,omitempty"`
}

func (o UpdateVgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVgwRequestBodyContent", string(data)}, " ")
}
