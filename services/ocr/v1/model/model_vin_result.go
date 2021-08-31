package model

import (
	"encoding/json"

	"strings"
)

//
type VinResult struct {
	// 识别检测到的车架号。

	Vin string `json:"vin"`
}

func (o VinResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VinResult struct{}"
	}

	return strings.Join([]string{"VinResult", string(data)}, " ")
}
