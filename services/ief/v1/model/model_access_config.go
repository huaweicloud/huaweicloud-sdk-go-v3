package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfig 访问配置，与access_protocol直接关联
type AccessConfig struct {
	ProtocolType *ValueInAttributes `json:"protocol_type"`

	ProtocolName *ValueInAttributes `json:"protocol_name,omitempty"`

	SlaveId *ValueInAttributes `json:"slave_id,omitempty"`

	Ip *ValueInAttributes `json:"ip,omitempty"`

	Port *ValueInAttributes `json:"port,omitempty"`

	SerialPort *ValueInAttributes `json:"serial_port,omitempty"`

	BaudRate *ValueInAttributes `json:"baud_rate,omitempty"`

	DataBits *ValueInAttributes `json:"data_bits,omitempty"`

	StopBits *ValueInAttributes `json:"stop_bits,omitempty"`

	ParityBits *ValueInAttributes `json:"parity_bits,omitempty"`

	Url *ValueInAttributes `json:"url,omitempty"`

	SecMode *ValueInAttributes `json:"sec_mode,omitempty"`

	SecPolicy *ValueInAttributes `json:"sec_policy,omitempty"`

	AuthType *ValueInAttributes `json:"auth_type,omitempty"`

	Username *ValueInAttributes `json:"username,omitempty"`

	Password *ValueInAttributes `json:"password,omitempty"`

	PrivateKey *ValueInAttributes `json:"private_key,omitempty"`

	Certificate *ValueInAttributes `json:"certificate,omitempty"`

	Timeout *ValueInAttributes `json:"timeout,omitempty"`
}

func (o AccessConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfig struct{}"
	}

	return strings.Join([]string{"AccessConfig", string(data)}, " ")
}
