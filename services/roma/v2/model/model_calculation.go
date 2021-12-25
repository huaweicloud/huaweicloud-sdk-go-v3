package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Calculation struct {
	// 从未连接过的设备数量

	NeverConnected *int32 `json:"never_connected,omitempty"`
	// 在线设备数量

	Online *int32 `json:"online,omitempty"`
	// 离线设备数量

	Offline *int32 `json:"offline,omitempty"`
}

func (o Calculation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Calculation struct{}"
	}

	return strings.Join([]string{"Calculation", string(data)}, " ")
}
