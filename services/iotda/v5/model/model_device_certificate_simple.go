package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceCertificateSimple 设备证书结构体
type DeviceCertificateSimple struct {

	// **参数说明**：设备证书ID。用来唯一标识一个设备证书
	CertificateId *string `json:"certificate_id,omitempty"`

	// **参数说明**：设备证书通用名称。
	CommonName *string `json:"common_name,omitempty"`

	// **参数说明**：设备证书过期时间。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// **参数说明**：设备证书SHA256指纹。
	Fingerprint *string `json:"fingerprint,omitempty"`

	// **参数说明**：设备证书状态，默认状态为激活 - ACTIVE：激活状态。 - INACTIVE：停用状态。
	Status *string `json:"status,omitempty"`
}

func (o DeviceCertificateSimple) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceCertificateSimple struct{}"
	}

	return strings.Join([]string{"DeviceCertificateSimple", string(data)}, " ")
}
