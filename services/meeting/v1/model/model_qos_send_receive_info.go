package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议与会者发送QoS数据。当qosType = audio/video/screen 时有效。
type QosSendReceiveInfo struct {

	// 码率, 单位kbps，不含阈值告警。当qosType = audio/video/screen 时有效。
	Bitrate *[]QosDataNoThrElement `json:"bitrate,omitempty"`

	// 时延，单位毫秒, 含阈值告警。当qosType = audio/video/screen 时有效。
	Latency *[]QosDataElement `json:"latency,omitempty"`

	// 抖动, 单位毫秒，含阈值告警。当qosType = audio/video/screen 时有效。
	Jitter *[]QosDataElement `json:"jitter,omitempty"`

	// 最大丢包率, 单位百分比 含阈值告警。当qosType = audio/video/screen 时有效。
	PacketLossMax *[]QosDataElement `json:"packet_loss_max,omitempty"`

	// 分辨率, 不含阈值告警。当qosType = video/screen 时有效。
	Resolution *[]QosDataNoThrElement `json:"resolution,omitempty"`

	// 帧率, 单位fps，不含阈值告警。当qosType = video/screen 时有效。
	Frame *[]QosDataNoThrElement `json:"frame,omitempty"`
}

func (o QosSendReceiveInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosSendReceiveInfo struct{}"
	}

	return strings.Join([]string{"QosSendReceiveInfo", string(data)}, " ")
}
