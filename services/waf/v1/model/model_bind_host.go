package model

import (
	"encoding/json"

	"strings"
)

// 绑定的域名信息
type BindHost struct {
	// 域名ID

	Id *string `json:"id,omitempty"`
	// 域名

	Hostname *string `json:"hostname,omitempty"`
	// 域名对应模式：cloud/premium

	WafType *string `json:"waf_type,omitempty"`
	// （仅独享模式）特殊域名模式

	Mode *string `json:"mode,omitempty"`
}

func (o BindHost) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BindHost struct{}"
	}

	return strings.Join([]string{"BindHost", string(data)}, " ")
}
