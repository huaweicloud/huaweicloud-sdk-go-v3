package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WordWaterMarkInfo struct {

	// 水印文字内容，必填 字幕内容1-64 type为1或2时必填，type为0时非必填
	Format string `json:"format"`

	// 参数校验：首位为#、除#外为6位或8位，每位字符由 0-9、a~f、A~F 组成【务必校验，错误时转码默认白色】
	FontColor *string `json:"font_color,omitempty"`

	// 字体大小
	FontSize *int32 `json:"font_size,omitempty"`

	// 字体，缺省值 空，可选值：harmonyRegular（鸿蒙）、douyu2.0（斗鱼追光体）
	Font *string `json:"font,omitempty"`

	// 时区，【取值参考 UTC-1200 至 UTC+1200，前三位UTC，第四位+或-，五六位表示小时，七八位固定00】
	TimeZone *string `json:"time_zone,omitempty"`

	// 缺省值 none，参数校验：首位为#、除#外为6位或8位，每位字符由 0-9、a~f、A~F 组成【务必校验，错误时转码默认黑色】
	ShadowColor *string `json:"shadow_color,omitempty"`

	Location *WatermarkLocation `json:"location"`
}

func (o WordWaterMarkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordWaterMarkInfo struct{}"
	}

	return strings.Join([]string{"WordWaterMarkInfo", string(data)}, " ")
}
