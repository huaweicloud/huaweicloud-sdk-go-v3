package model

import (
	"encoding/json"

	"strings"
)

// 弹性IP信息
type PublicipResult struct {
	Eip *EipResult `json:"eip,omitempty"`
}

func (o PublicipResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublicipResult struct{}"
	}

	return strings.Join([]string{"PublicipResult", string(data)}, " ")
}
