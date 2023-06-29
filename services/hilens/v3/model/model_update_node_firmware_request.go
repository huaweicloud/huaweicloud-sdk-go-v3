package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeFirmwareRequest Request Object
type UpdateNodeFirmwareRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`

	// 固件ID,，从专业版HiLens控制台固件管理[查询固件列表](ListFirmwares.xml)获取
	FirmwareId string `json:"firmware_id"`

	// 离线场景超期时间，单位分钟，范围在1-86400
	XExpiredTime *int32 `json:"X-Expired-Time,omitempty"`
}

func (o UpdateNodeFirmwareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeFirmwareRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodeFirmwareRequest", string(data)}, " ")
}
