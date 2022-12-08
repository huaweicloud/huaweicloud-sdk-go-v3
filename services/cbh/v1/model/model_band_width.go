package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 带宽参数
type BandWidth struct {

	// 带宽，按需1-100，包年包月1-300
	Size string `json:"size"`

	// 共享类型 目前只支持PER
	ShareType string `json:"share_type"`

	// 带宽计费类型 未traffic或为空
	ChargeMode string `json:"charge_mode"`

	// 带宽对应产品ID
	ProductId string `json:"product_id"`
}

func (o BandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandWidth struct{}"
	}

	return strings.Join([]string{"BandWidth", string(data)}, " ")
}
