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

//
type ExtraDhcpOpt struct {
	// Option名称
	OptName *string `json:"opt_name,omitempty"`
	// Option值
	OptValue *string `json:"opt_value,omitempty"`
}

func (o ExtraDhcpOpt) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ExtraDhcpOpt", string(data)}, " ")
}
