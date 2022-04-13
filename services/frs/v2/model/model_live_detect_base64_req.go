package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LiveDetectBase64Req struct {
	// 该参数为动作时间数组拼接的字符串，数组的长度和actions的数量一致，每一项代表了对应次序动作的起始时间和结束时间，单位为距视频开始的毫秒数。

	ActionTime *string `json:"action_time,omitempty"`
	// 视频数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议客户端压缩到200KB~2MB。 • 限制视频时长1～15秒。 • 建议帧率10fps～30fps。 • 封装格式：mp4、avi、flv、webm、asf、mov。 • 视频编码格式： h261、h263、h264、hevc、vc1、vp8、vp9、wmv3。

	VideoBase64 string `json:"video_base64"`
	// 动作代码顺序列表，英文逗号（,）分隔。建议单动作，目前支持的动作有： • 1：左摇头 • 2：右摇头 • 3：点头 • 4：嘴部动作

	Actions string `json:"actions"`
}

func (o LiveDetectBase64Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveDetectBase64Req struct{}"
	}

	return strings.Join([]string{"LiveDetectBase64Req", string(data)}, " ")
}
