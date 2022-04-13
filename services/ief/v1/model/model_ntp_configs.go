package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ntp配置
type NtpConfigs struct {
	// ntp服务是否开启

	NtpEnabled *bool `json:"ntp_enabled,omitempty"`
	// ntp server地址

	Ntpservers *[]string `json:"ntpservers,omitempty"`
}

func (o NtpConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NtpConfigs struct{}"
	}

	return strings.Join([]string{"NtpConfigs", string(data)}, " ")
}
