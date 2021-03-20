package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoInfo struct {
	// 视频宽度

	Width *int32 `json:"width,omitempty"`
	// 视频高度

	Height *int32 `json:"height,omitempty"`
	// 视频码率，单位: kbit/s

	Bitrate *int32 `json:"bitrate,omitempty"`
	// 视频码率，单位: bit/s

	BitrateBps *int64 `json:"bitrate_bps,omitempty"`
	// 帧率。    取值范围：0或[5,60]，0表示自适应。    单位：帧每秒。    > 若设置的帧率不在取值范围内，则自动调整为0，若设置的帧率高于片源帧率，则自动调整为片源帧率。

	FrameRate *int32 `json:"frame_rate,omitempty"`
	// 视频编码格式

	Codec *string `json:"codec,omitempty"`
	// 片源动态范围类型。  取值如下： - SDR - HDR10 - CUVA_HDR

	DynamicRange *VideoInfoDynamicRange `json:"dynamic_range,omitempty"`
}

func (o VideoInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VideoInfo struct{}"
	}

	return strings.Join([]string{"VideoInfo", string(data)}, " ")
}

type VideoInfoDynamicRange struct {
	value string
}

type VideoInfoDynamicRangeEnum struct {
	SDR      VideoInfoDynamicRange
	HDR10    VideoInfoDynamicRange
	CUVA_HDR VideoInfoDynamicRange
}

func GetVideoInfoDynamicRangeEnum() VideoInfoDynamicRangeEnum {
	return VideoInfoDynamicRangeEnum{
		SDR: VideoInfoDynamicRange{
			value: "SDR",
		},
		HDR10: VideoInfoDynamicRange{
			value: "HDR10",
		},
		CUVA_HDR: VideoInfoDynamicRange{
			value: "CUVA_HDR",
		},
	}
}

func (c VideoInfoDynamicRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VideoInfoDynamicRange) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
