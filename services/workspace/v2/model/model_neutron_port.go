package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronPort 端口详情。
type NeutronPort struct {

	// 端口唯一标识。
	Id *string `json:"id,omitempty"`

	// 私有ip状态。
	Status *string `json:"status,omitempty"`

	// 端口所属网络的ID。
	NetworkId *string `json:"network_id,omitempty"`

	// 端口所属设备的Id。
	DeviceId *string `json:"device_id,omitempty"`

	// 端口IP。
	FixedIps *[]FixedIp `json:"fixed_ips,omitempty"`
}

func (o NeutronPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronPort struct{}"
	}

	return strings.Join([]string{"NeutronPort", string(data)}, " ")
}
