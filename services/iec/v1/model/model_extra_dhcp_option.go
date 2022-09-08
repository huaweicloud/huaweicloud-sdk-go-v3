package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DHCP扩展属性
type ExtraDhcpOption struct {

	// Option名称
	OptName *string `json:"opt_name,omitempty"`

	// Option值
	OptValue *string `json:"opt_value,omitempty"`
}

func (o ExtraDhcpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraDhcpOption struct{}"
	}

	return strings.Join([]string{"ExtraDhcpOption", string(data)}, " ")
}
