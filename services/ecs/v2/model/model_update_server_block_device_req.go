package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerBlockDeviceReq This is a auto create Body Object
type UpdateServerBlockDeviceReq struct {
	BlockDevice *UpdateServerBlockDeviceOption `json:"block_device"`
}

func (o UpdateServerBlockDeviceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerBlockDeviceReq struct{}"
	}

	return strings.Join([]string{"UpdateServerBlockDeviceReq", string(data)}, " ")
}
