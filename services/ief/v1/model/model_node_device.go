package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点的终端设备信息
type NodeDevice struct {
	Added *DeviceInfos `json:"added,omitempty" xml:"added"`

	// 要解绑的终端设备ID
	Removed *[]string `json:"removed,omitempty" xml:"removed"`
}

func (o NodeDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDevice struct{}"
	}

	return strings.Join([]string{"NodeDevice", string(data)}, " ")
}
