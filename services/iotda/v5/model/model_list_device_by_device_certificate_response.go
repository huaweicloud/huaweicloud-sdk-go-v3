package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceByDeviceCertificateResponse Response Object
type ListDeviceByDeviceCertificateResponse struct {

	// 设备证书列表
	Devices *[]DeviceSimple `json:"devices,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListDeviceByDeviceCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceByDeviceCertificateResponse struct{}"
	}

	return strings.Join([]string{"ListDeviceByDeviceCertificateResponse", string(data)}, " ")
}
