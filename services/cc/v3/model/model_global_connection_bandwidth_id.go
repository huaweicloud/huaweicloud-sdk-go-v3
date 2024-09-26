package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalConnectionBandwidthId 全域互联带宽ID。
type GlobalConnectionBandwidthId struct {

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`
}

func (o GlobalConnectionBandwidthId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthId struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthId", string(data)}, " ")
}
