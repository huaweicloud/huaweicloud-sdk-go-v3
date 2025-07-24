package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerPowerStatus 服务器电源状态对象
type ServerPowerStatus struct {

	// 开机状态的数量
	On int32 `json:"on"`

	// 关机状态的数量
	Off int32 `json:"off"`

	// 未知状态的数量
	Unknown *int32 `json:"unknown,omitempty"`
}

func (o ServerPowerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerPowerStatus struct{}"
	}

	return strings.Join([]string{"ServerPowerStatus", string(data)}, " ")
}
