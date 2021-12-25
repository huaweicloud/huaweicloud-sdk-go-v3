package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 访问配置，与access_protocol直接关联
type AccessConfig struct {
	ProtocolType *AccessConfigProtocolType `json:"protocol_type"`

	ProtocolName *AccessConfigProtocolName `json:"protocol_name,omitempty"`

	SlaveId *AccessConfigSlaveId `json:"slave_id,omitempty"`

	Ip *AccessConfigIp `json:"ip,omitempty"`

	Port *AccessConfigPort `json:"port,omitempty"`

	SerialPort *AccessConfigSerialPort `json:"serial_port,omitempty"`

	BaudRate *AccessConfigBaudRate `json:"baud_rate,omitempty"`

	DataBits *AccessConfigDataBits `json:"data_bits,omitempty"`

	StopBits *AccessConfigStopBits `json:"stop_bits,omitempty"`

	ParityBits *AccessConfigParityBits `json:"parity_bits,omitempty"`

	Url *AccessConfigUrl `json:"url,omitempty"`

	SecMode *AccessConfigSecMode `json:"sec_mode,omitempty"`

	SecPolicy *AccessConfigSecPolicy `json:"sec_policy,omitempty"`

	AuthType *AccessConfigAuthType `json:"auth_type,omitempty"`

	Username *AccessConfigUsername `json:"username,omitempty"`

	Password *AccessConfigPassword `json:"password,omitempty"`

	PrivateKey *AccessConfigPrivateKey `json:"private_key,omitempty"`

	Certificate *AccessConfigCertificate `json:"certificate,omitempty"`

	Timeout *AccessConfigTimeout `json:"timeout,omitempty"`
}

func (o AccessConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfig struct{}"
	}

	return strings.Join([]string{"AccessConfig", string(data)}, " ")
}
