package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryStreamInfo struct {

	// 推流域名。  - type为0表主播推流域名。  - type为1表示第三方推流域名
	Domain *string `json:"domain,omitempty" xml:"domain"`

	// 应用名称。
	App *string `json:"app,omitempty" xml:"app"`

	// 流名。
	Stream *string `json:"stream,omitempty" xml:"stream"`

	// 推流类型，取值如下：  - 0：表示主播推流  - 1：表示第三方推流
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 视频编码格式。
	VideoCodec *string `json:"video_codec,omitempty" xml:"video_codec"`

	// 音频编码格式。
	AudioCodec *string `json:"audio_codec,omitempty" xml:"audio_codec"`

	// 主播ip。
	ClientIp *string `json:"client_ip,omitempty" xml:"client_ip"`

	// 采样开始时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 采样结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o HistoryStreamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryStreamInfo struct{}"
	}

	return strings.Join([]string{"HistoryStreamInfo", string(data)}, " ")
}
