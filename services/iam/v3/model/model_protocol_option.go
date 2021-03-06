package model

import (
	"encoding/json"

	"strings"
)

//
type ProtocolOption struct {
	// 映射ID。

	MappingId string `json:"mapping_id"`
}

func (o ProtocolOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProtocolOption struct{}"
	}

	return strings.Join([]string{"ProtocolOption", string(data)}, " ")
}
