/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 子网配置的NTP地址对象
type ExtraDhcpOption struct {
	// 功能说明：子网配置的NTP地址名称 约束：目前只支持字段“ntp”
	OptName ExtraDhcpOptionOptName `json:"opt_name"`
	// 功能说明：子网配置的NTP地址 约束：目前只支持IPv4地址，每个IP地址以逗号隔开，IP地址个数不能超过4个，不能存在相同地址。该字段为null表示取消该子网NTP的设置，不能为””(空字符串)。
	OptValue *string `json:"opt_value,omitempty"`
}

func (o ExtraDhcpOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ExtraDhcpOption", string(data)}, " ")
}

type ExtraDhcpOptionOptName struct {
	value string
}

type ExtraDhcpOptionOptNameEnum struct {
	NTP ExtraDhcpOptionOptName
}

func GetExtraDhcpOptionOptNameEnum() ExtraDhcpOptionOptNameEnum {
	return ExtraDhcpOptionOptNameEnum{
		NTP: ExtraDhcpOptionOptName{
			value: "ntp",
		},
	}
}

func (c ExtraDhcpOptionOptName) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ExtraDhcpOptionOptName) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
