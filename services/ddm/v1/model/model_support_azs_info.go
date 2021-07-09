package model

import (
	"encoding/json"

	"strings"
)

type SupportAzsInfo struct {
	// 可用区编码。

	Code *string `json:"code,omitempty"`
	// 可用区名称。

	Name *string `json:"name,omitempty"`
	// 是否支持。

	Favored *bool `json:"favored,omitempty"`
}

func (o SupportAzsInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SupportAzsInfo struct{}"
	}

	return strings.Join([]string{"SupportAzsInfo", string(data)}, " ")
}
