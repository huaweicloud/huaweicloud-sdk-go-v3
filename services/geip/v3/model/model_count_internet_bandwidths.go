package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountInternetBandwidths struct {

	// 全域公网带宽的数目
	Count int32 `json:"count"`
}

func (o CountInternetBandwidths) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInternetBandwidths struct{}"
	}

	return strings.Join([]string{"CountInternetBandwidths", string(data)}, " ")
}
