package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Qos 网络质量对象，音频不涉及分辨率和帧率（注：所有媒体流的数据均为服务器视角，Portal展示时需要转换为端侧视角）
type Qos struct {

	// 网络质量评级
	NetworkQuality *string `json:"networkQuality,omitempty"`

	// 编解码类型
	CodecType *string `json:"codecType,omitempty"`

	// 带宽(kbit/s)
	BandWidth *int32 `json:"bandWidth,omitempty"`

	// 丢包率（千分数）
	LostPacketRate *int32 `json:"lostPacketRate,omitempty"`

	// 时延(ms)
	Delay *int32 `json:"delay,omitempty"`

	// 抖动(ms)
	Jitter *int32 `json:"jitter,omitempty"`

	// 分辨率:高
	ResolutionHeight *int32 `json:"resolutionHeight,omitempty"`

	// 分辨率：宽
	ResolutionWidth *int32 `json:"resolutionWidth,omitempty"`

	// 帧率
	FrameRate *int32 `json:"frameRate,omitempty"`

	// 该媒体流编码会场ID，仅服务器向端侧发送方向涉及(视频)，如路径A-->服务器-->B，即B观看A，该媒体流的codec_user_id为A
	CodecUserId *string `json:"codecUserId,omitempty"`
}

func (o Qos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Qos struct{}"
	}

	return strings.Join([]string{"Qos", string(data)}, " ")
}
