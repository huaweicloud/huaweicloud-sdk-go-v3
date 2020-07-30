/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type ExtraDhcpOpt struct {
	// Option名称
	OptName string `json:"opt_name,omitempty"`
	// Option值
	OptValue string `json:"opt_value,omitempty"`
}
