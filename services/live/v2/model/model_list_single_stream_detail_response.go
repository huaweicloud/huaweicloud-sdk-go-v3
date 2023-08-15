package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSingleStreamDetailResponse Response Object
type ListSingleStreamDetailResponse struct {

	// 推流域名
	PublishDomain *string `json:"publish_domain,omitempty"`

	// 应用名
	App *string `json:"app,omitempty"`

	// 流名
	Stream *string `json:"stream,omitempty"`

	// 展示流视频帧率情况，帧率单位为fps。  如果出现断流则会出现多个时间段流信息，如： ``` \"video_framerate\": [     {       \"start_time\": \"2022-02-04T07:00:00Z\",       \"end_time\": \"2022-02-04T07:00:02Z\",       \"data_list\": [         21,         22       ]     },     {       \"start_time\": \"2022-02-04T07:00:05Z\",       \"end_time\": \"2022-02-04T07:00:07Z\",       \"data_list\": [         13,         34,         21       ]     }   ]
	VideoFramerate *[]StreamDetail `json:"video_framerate,omitempty"`

	// 展示流视频码率情况，码率单位为Kbps。  如果出现断流则会出现多个时间段流信息，如： ``` \"video_bitrate\": [     {       \"start_time\": \"2022-02-04T07:00:00Z\",       \"end_time\": \"2022-02-04T07:00:02Z\",       \"data_list\": [         1326,         1268,         775       ]     },     {       \"start_time\": \"2022-02-04T07:00:05Z\",       \"end_time\": \"2022-02-04T07:00:07Z\",       \"data_list\": [         1021,         2022       ]     }   ]
	VideoBitrate *[]StreamDetail `json:"video_bitrate,omitempty"`

	// 展示流音频帧率情况，帧率单位为fps。  如果出现断流则会出现多个时间段流信息，如： ``` \"audio_framerate\": [     {       \"start_time\": \"2022-02-04T07:00:00Z\",       \"end_time\": \"2022-02-04T07:00:02Z\",       \"data_list\": [         10,         17       ]     },     {       \"start_time\": \"2022-02-04T07:00:05Z\",       \"end_time\": \"2022-02-04T07:00:06Z\",       \"data_list\": [         31,         33       ]     }   ]
	AudioFramerate *[]StreamDetail `json:"audio_framerate,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSingleStreamDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSingleStreamDetailResponse struct{}"
	}

	return strings.Join([]string{"ListSingleStreamDetailResponse", string(data)}, " ")
}
