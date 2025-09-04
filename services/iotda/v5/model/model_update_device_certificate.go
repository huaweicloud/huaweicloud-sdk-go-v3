package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceCertificate 更新设备证书结构体。
type UpdateDeviceCertificate struct {

	// **参数说明**：设备证书状态，默认状态为激活 - ACTIVE：激活状态。 - INACTIVE：停用状态。
	Status string `json:"status"`
}

func (o UpdateDeviceCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceCertificate struct{}"
	}

	return strings.Join([]string{"UpdateDeviceCertificate", string(data)}, " ")
}
