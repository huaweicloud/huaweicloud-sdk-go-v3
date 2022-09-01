package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ntp配置
type NtpConfigs struct {

	// ntp服务是否开启
	NtpEnabled *bool `json:"ntp_enabled,omitempty" xml:"ntp_enabled"`

	// ntp server地址
	Ntpservers *[]string `json:"ntpservers,omitempty" xml:"ntpservers"`
}

func (o NtpConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NtpConfigs struct{}"
	}

	return strings.Join([]string{"NtpConfigs", string(data)}, " ")
}
