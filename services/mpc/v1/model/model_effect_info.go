package model

import (
	"encoding/json"

	"strings"
)

type EffectInfo struct {
	// 特效处理开始时间点，单位：秒，精确到小数点后4位

	StartTime *string `json:"start_time,omitempty"`
	// 特效处理结束时间点，单位：秒，精确到小数点后4位

	StopTime *string `json:"stop_time,omitempty"`
	// 相对视频源左上角顶点的水平偏移位置坐标。整数型表示像素值，范围[0,4096]，单位px

	Dx *string `json:"dx,omitempty"`
	// 相对视频源左上角顶点的垂直偏移位置坐标。整数型表示像素值，范围[0,4096]，单位px

	Dy *string `json:"dy,omitempty"`
	// 特效区域的宽。整数型表示像素值，范围(0,4096]，单位px

	Width *string `json:"width,omitempty"`
	// 特效区域的高。整数型表示像素值，范围(0,4096]，单位px

	Height *string `json:"height,omitempty"`
}

func (o EffectInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EffectInfo struct{}"
	}

	return strings.Join([]string{"EffectInfo", string(data)}, " ")
}
