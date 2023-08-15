package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeDevice 边缘节点的终端设备信息
type NodeDevice struct {
	Added *DeviceInfos `json:"added,omitempty"`

	// 要解绑的终端设备ID
	Removed *[]string `json:"removed,omitempty"`
}

func (o NodeDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDevice struct{}"
	}

	return strings.Join([]string{"NodeDevice", string(data)}, " ")
}
