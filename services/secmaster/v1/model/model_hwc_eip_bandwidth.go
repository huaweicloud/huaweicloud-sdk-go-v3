package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEipBandwidth 带宽
type HwcEipBandwidth struct {

	// 带宽ID
	Id *string `json:"id,omitempty"`

	// 带宽大小 取值范围：默认5Mbit/s~2000Mbit/s
	Size *int32 `json:"size,omitempty"`

	// 带宽类型，标识是否是共享带宽 取值范围： PER：独享带宽 WHOLE：共享带宽 约束：其中IPv6暂不支持WHOLE类型带宽。
	ShareType *string `json:"share_type,omitempty"`

	// 带宽名称
	Name *string `json:"name,omitempty"`
}

func (o HwcEipBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEipBandwidth struct{}"
	}

	return strings.Join([]string{"HwcEipBandwidth", string(data)}, " ")
}
