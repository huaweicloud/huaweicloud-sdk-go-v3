package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerBlockDeviceRequest Request Object
type UpdateServerBlockDeviceRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	// 磁盘id，uuid格式
	VolumeId string `json:"volume_id"`

	Body *UpdateServerBlockDeviceReq `json:"body,omitempty"`
}

func (o UpdateServerBlockDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerBlockDeviceRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerBlockDeviceRequest", string(data)}, " ")
}
