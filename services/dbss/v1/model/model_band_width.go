package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandWidth 带宽参数
type BandWidth struct {

	// 计费类型
	Chargemode *string `json:"chargemode,omitempty"`

	// 带宽对应产品ID
	Productid *string `json:"productid,omitempty"`

	// 共享类型，目前仅支持PER
	Sharetype *string `json:"sharetype,omitempty"`

	// 带宽，按需1-100M，包年包月1-300M
	Size *int32 `json:"size,omitempty"`
}

func (o BandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandWidth struct{}"
	}

	return strings.Join([]string{"BandWidth", string(data)}, " ")
}
