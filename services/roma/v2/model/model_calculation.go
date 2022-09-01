package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Calculation struct {

	// 从未连接过的设备数量
	NeverConnected *int32 `json:"never_connected,omitempty" xml:"never_connected"`

	// 在线设备数量
	Online *int32 `json:"online,omitempty" xml:"online"`

	// 离线设备数量
	Offline *int32 `json:"offline,omitempty" xml:"offline"`
}

func (o Calculation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Calculation struct{}"
	}

	return strings.Join([]string{"Calculation", string(data)}, " ")
}
