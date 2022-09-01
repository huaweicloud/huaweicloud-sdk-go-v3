package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkTrafficStats struct {

	// 下行吞吐量（byte）
	OutputThroughput *int64 `json:"output_throughput,omitempty" xml:"output_throughput"`

	// 上行吞吐量（byte）
	InputThroughput *int64 `json:"input_throughput,omitempty" xml:"input_throughput"`
}

func (o NetworkTrafficStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkTrafficStats struct{}"
	}

	return strings.Join([]string{"NetworkTrafficStats", string(data)}, " ")
}
