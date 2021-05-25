package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBandwidthRequest struct {
	// 带宽唯一标识

	BandwidthId string `json:"bandwidth_id"`
}

func (o ShowBandwidthRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ShowBandwidthRequest", string(data)}, " ")
}
