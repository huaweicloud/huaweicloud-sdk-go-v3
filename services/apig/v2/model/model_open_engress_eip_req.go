package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenEngressEipReq struct {
	// 出公网带宽  单位：Mbit/s

	BandwidthSize *string `json:"bandwidth_size,omitempty"`
}

func (o OpenEngressEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenEngressEipReq struct{}"
	}

	return strings.Join([]string{"OpenEngressEipReq", string(data)}, " ")
}
