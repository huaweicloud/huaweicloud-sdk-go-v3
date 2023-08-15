package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NtpConfigs ntp配置
type NtpConfigs struct {

	// ntp服务是否开启
	NtpEnabled *bool `json:"ntp_enabled,omitempty"`

	// ntp server地址
	NtpServers *[]string `json:"ntp_servers,omitempty"`
}

func (o NtpConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NtpConfigs struct{}"
	}

	return strings.Join([]string{"NtpConfigs", string(data)}, " ")
}
