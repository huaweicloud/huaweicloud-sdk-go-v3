package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCarouselTaskDetailResponse Response Object
type ListCarouselTaskDetailResponse struct {

	// 采样开始时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime *string `json:"start_time,omitempty"`

	// 采样结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"end_time,omitempty"`

	// 视频帧率信息列表，单位为fps。
	VideoFramerateList *[]int64 `json:"video_framerate_list,omitempty"`

	// 视频码率信息列表，单位为kbps。
	VideoBitrateList *[]int64 `json:"video_bitrate_list,omitempty"`

	// 音频帧率信息列表，单位为fps。
	AudioFramerateList *[]int64 `json:"audio_framerate_list,omitempty"`

	// 音频码率信息列表，单位为kbps。
	AudioBitrateList *[]int64 `json:"audio_bitrate_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCarouselTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCarouselTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ListCarouselTaskDetailResponse", string(data)}, " ")
}
