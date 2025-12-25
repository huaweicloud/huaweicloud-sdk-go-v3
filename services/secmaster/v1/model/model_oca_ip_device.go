package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIpDevice 设备信息
type OcaIpDevice struct {

	// 设备类型
	Type *string `json:"type,omitempty"`

	// 设备型号
	Model *string `json:"model,omitempty"`

	Vendor *OcaIpVendor `json:"vendor,omitempty"`
}

func (o OcaIpDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIpDevice struct{}"
	}

	return strings.Join([]string{"OcaIpDevice", string(data)}, " ")
}
