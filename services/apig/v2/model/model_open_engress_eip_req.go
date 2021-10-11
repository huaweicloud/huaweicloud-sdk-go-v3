package model

import (
	"encoding/json"

	"strings"
)

type OpenEngressEipReq struct {
	// 出公网带宽  单位：Mbit/s

	BandwidthSize *string `json:"bandwidth_size,omitempty"`
}

func (o OpenEngressEipReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpenEngressEipReq struct{}"
	}

	return strings.Join([]string{"OpenEngressEipReq", string(data)}, " ")
}
