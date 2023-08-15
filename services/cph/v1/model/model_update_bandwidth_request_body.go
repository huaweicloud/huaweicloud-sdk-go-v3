package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBandwidthRequestBody struct {

	// - 小于等于300Mbit/s：默认最小增长步长为1Mbit/s。 - 300Mbit/s~1000Mbit/s：默认最小增长步长为50Mbit/s。 - 大于1000Mbit/s：默认最小增长步长为500Mbit/s。
	BandWidthSize int32 `json:"band_width_size"`
}

func (o UpdateBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthRequestBody", string(data)}, " ")
}
