package model

import (
	"encoding/json"

	"strings"
)

type AvailableZone struct {
	// 可用区名称。

	Name *string `json:"name,omitempty"`
	// 实例创建失败错误信息

	Id *string `json:"id,omitempty"`
	// 可用区编码。

	Code *string `json:"code,omitempty"`
	// 可用区端口号。

	Port *string `json:"port,omitempty"`

	LocalName *LocalName `json:"local_name,omitempty"`
	// 可用区支持的实例规格。

	Specs map[string]bool `json:"specs,omitempty"`
}

func (o AvailableZone) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AvailableZone struct{}"
	}

	return strings.Join([]string{"AvailableZone", string(data)}, " ")
}
