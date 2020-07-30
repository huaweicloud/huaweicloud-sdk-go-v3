/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 子网配置的NTP地址对象
type ExtraDhcpOption struct {
	// 功能说明：子网配置的NTP地址名称 约束：目前只支持字段“ntp”
	OptName string `json:"opt_name"`
	// 功能说明：子网配置的NTP地址 约束：目前只支持IPv4地址，每个IP地址以逗号隔开，IP地址个数不能超过4个，不能存在相同地址。该字段为null表示取消该子网NTP的设置，不能为””(空字符串)。
	OptValue string `json:"opt_value,omitempty"`
}
