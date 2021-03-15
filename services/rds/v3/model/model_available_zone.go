package model

import (
	"encoding/json"

	"strings"
)

type AvailableZone struct {
	// 可用区编码。
	Code string `json:"code"`
	// 可用区描述。
	Description string `json:"description"`
}

func (o AvailableZone) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AvailableZone struct{}"
	}

	return strings.Join([]string{"AvailableZone", string(data)}, " ")
}
