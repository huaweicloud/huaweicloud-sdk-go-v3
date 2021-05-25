package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSharedBandwidthRequest struct {
	// 带宽唯一标识  约束： 当前仅支持删除共享带宽

	BandwidthId string `json:"bandwidth_id"`
}

func (o DeleteSharedBandwidthRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSharedBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DeleteSharedBandwidthRequest", string(data)}, " ")
}
