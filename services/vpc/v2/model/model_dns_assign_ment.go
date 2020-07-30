/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type DnsAssignMent struct {
	// 端口hostname
	Hostname string `json:"hostname,omitempty"`
	// 端口IP地址
	IpAddress string `json:"ip_address,omitempty"`
	// 端口内网fqdn
	Fqdn string `json:"fqdn,omitempty"`
}
