package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoCoverAnalysisinference struct {

	// 动态封面帧率分子
	GifRatenum *int32 `json:"gif_ratenum,omitempty"`

	// 动态封面帧率分母
	GifRateden *int32 `json:"gif_rateden,omitempty"`

	// 动态封面长边长度
	GifLongSide *int32 `json:"gif_long_side,omitempty"`

	// 动态封面短边长度
	GifShortSide *int32 `json:"gif_short_side,omitempty"`
}

func (o VideoCoverAnalysisinference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCoverAnalysisinference struct{}"
	}

	return strings.Join([]string{"VideoCoverAnalysisinference", string(data)}, " ")
}
