package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InternetBandwidthInfo struct {

	// 全域公网带宽的ID
	Id *string `json:"id,omitempty"`

	// 全域公网带宽大小（出云方向）
	Size *int32 `json:"size,omitempty"`
}

func (o InternetBandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InternetBandwidthInfo struct{}"
	}

	return strings.Join([]string{"InternetBandwidthInfo", string(data)}, " ")
}
