package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgemgrDevicesUpdate 终端设备更新配置
type EdgemgrDevicesUpdate struct {
	Device *EdgemgrDevicesPara `json:"device"`
}

func (o EdgemgrDevicesUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgemgrDevicesUpdate struct{}"
	}

	return strings.Join([]string{"EdgemgrDevicesUpdate", string(data)}, " ")
}
