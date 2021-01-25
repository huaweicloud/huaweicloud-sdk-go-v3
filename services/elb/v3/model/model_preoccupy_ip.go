package model

import (
	"encoding/json"

	"strings"
)

//
type PreoccupyIp struct {
	// 预占IP总数
	Total int32 `json:"total"`
}

func (o PreoccupyIp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PreoccupyIp struct{}"
	}

	return strings.Join([]string{"PreoccupyIp", string(data)}, " ")
}
