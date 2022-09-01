package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 访问配置，与access_protocol直接关联
type AccessConfig struct {
	ProtocolType *ValueInAttributes `json:"protocol_type" xml:"protocol_type"`

	ProtocolName *ValueInAttributes `json:"protocol_name,omitempty" xml:"protocol_name"`

	SlaveId *ValueInAttributes `json:"slave_id,omitempty" xml:"slave_id"`

	Ip *ValueInAttributes `json:"ip,omitempty" xml:"ip"`

	Port *ValueInAttributes `json:"port,omitempty" xml:"port"`

	SerialPort *ValueInAttributes `json:"serial_port,omitempty" xml:"serial_port"`

	BaudRate *ValueInAttributes `json:"baud_rate,omitempty" xml:"baud_rate"`

	DataBits *ValueInAttributes `json:"data_bits,omitempty" xml:"data_bits"`

	StopBits *ValueInAttributes `json:"stop_bits,omitempty" xml:"stop_bits"`

	ParityBits *ValueInAttributes `json:"parity_bits,omitempty" xml:"parity_bits"`

	Url *ValueInAttributes `json:"url,omitempty" xml:"url"`

	SecMode *ValueInAttributes `json:"sec_mode,omitempty" xml:"sec_mode"`

	SecPolicy *ValueInAttributes `json:"sec_policy,omitempty" xml:"sec_policy"`

	AuthType *ValueInAttributes `json:"auth_type,omitempty" xml:"auth_type"`

	Username *ValueInAttributes `json:"username,omitempty" xml:"username"`

	Password *ValueInAttributes `json:"password,omitempty" xml:"password"`

	PrivateKey *ValueInAttributes `json:"private_key,omitempty" xml:"private_key"`

	Certificate *ValueInAttributes `json:"certificate,omitempty" xml:"certificate"`

	Timeout *ValueInAttributes `json:"timeout,omitempty" xml:"timeout"`
}

func (o AccessConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfig struct{}"
	}

	return strings.Join([]string{"AccessConfig", string(data)}, " ")
}
