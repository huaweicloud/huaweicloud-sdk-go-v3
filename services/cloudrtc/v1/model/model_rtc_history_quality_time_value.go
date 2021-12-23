package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RtcHistoryQualityTimeValue struct {
	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为YYYY-MM-DD

	Date *string `json:"date,omitempty"`
	// 加入房间成功率参数取值，取值为1代表成功率100%

	JoinSuccessRate *float64 `json:"join_success_rate,omitempty"`
	// 5s内加入房间成功率参数取值，取值为1代表成功率100%

	JoinSuccessIn5secsRate *float64 `json:"join_success_in5secs_rate,omitempty"`
	// 视频卡顿率参数取值，取值为1代表成功率100%

	VideoFreezeRate *float64 `json:"video_freeze_rate,omitempty"`
	// 音频卡顿率参数取值，取值为1代表成功率100%

	AudioFreezeRate *float64 `json:"audio_freeze_rate,omitempty"`
	// 首帧视频接收耗时，单位毫秒

	FirstVideoRecvTime *int64 `json:"first_video_recv_time,omitempty"`
	// 首帧音频接收耗时，单位毫秒

	FirstAudioRecvTime *int64 `json:"first_audio_recv_time,omitempty"`
	// 拉流成功率参数取值，取值为1代表成功率100%

	PullStreamSuccessRate *float64 `json:"pull_stream_success_rate,omitempty"`
	// 推流成功率参数取值，取值为1代表成功率100%

	PushStreamSuccessRate *float64 `json:"push_stream_success_rate,omitempty"`
	// 客户端视频上行优质传输率，取值为1代表成功率100%

	VideoUpstreamExcellentTransRate *float64 `json:"video_upstream_excellent_trans_rate,omitempty"`
	// 客户端音频上行优质传输率，取值为1代表成功率100%

	AudioUpstreamExcellentTransRate *float64 `json:"audio_upstream_excellent_trans_rate,omitempty"`
	// 端到端视频优质传输率，取值为1代表成功率100%

	VideoExcellentTransRate *float64 `json:"video_excellent_trans_rate,omitempty"`
	// 端到端音频优质传输率，取值为1代表成功率100%

	AudioExcellentTransRate *float64 `json:"audio_excellent_trans_rate,omitempty"`
	// 端到端视频网络时，单位为毫秒，取当天所有用户网络延迟的中位数，取值为1代表成功率100%

	VideoTransDelay *float64 `json:"video_trans_delay,omitempty"`
	// 端到端音频网络时延，单位为毫秒，取当天所有用户网络延迟的中位数，取值为1代表成功率100%

	AudioTransDelay *float64 `json:"audio_trans_delay,omitempty"`
}

func (o RtcHistoryQualityTimeValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcHistoryQualityTimeValue struct{}"
	}

	return strings.Join([]string{"RtcHistoryQualityTimeValue", string(data)}, " ")
}
