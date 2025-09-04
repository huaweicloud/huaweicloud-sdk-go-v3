package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceCertificateResponse Response Object
type ListDeviceCertificateResponse struct {

	// 设备证书列表
	DeviceCertificates *[]DeviceCertificateSimple `json:"device_certificates,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListDeviceCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceCertificateResponse struct{}"
	}

	return strings.Join([]string{"ListDeviceCertificateResponse", string(data)}, " ")
}
