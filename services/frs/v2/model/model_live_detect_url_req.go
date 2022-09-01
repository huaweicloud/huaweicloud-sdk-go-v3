package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LiveDetectUrlReq struct {

	// 该参数为动作时间数组拼接的字符串，数组的长度和actions的数量一致，每一项代表了对应次序动作的起始时间和结束时间，单位为距视频开始的毫秒数。
	ActionTime *string `json:"action_time,omitempty" xml:"action_time"`

	// [视频的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。 开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/api-face/face_02_0006.html)。视频要求： • 视频Base64编码后大小不超过8MB。 • 限制视频时长1～15秒。 • 建议帧率10fps～30fps。 • 封装格式：mp4、avi、flv、webm、asf、mov。 • 视频编码格式： h261、h263、h264、hevc、vc1、vp8、vp9、wmv3。](tag:hc) [视频的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。           开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0006.html)。视频要求：           • 视频Base64编码后大小不超过8MB。           • 限制视频时长1～15秒。           • 建议帧率10fps～30fps。           • 封装格式：mp4、avi、flv、webm、asf、mov。           • 视频编码格式： h261、h263、h264、hevc、vc1、vp8、vp9、wmv3。](tag:hk)
	VideoUrl string `json:"video_url" xml:"video_url"`

	// 动作代码顺序列表，英文逗号（,）分隔。建议单动作，目前支持的动作有： • 1：左摇头 • 2：右摇头 • 3：点头 • 4：嘴部动作
	Actions string `json:"actions" xml:"actions"`
}

func (o LiveDetectUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveDetectUrlReq struct{}"
	}

	return strings.Join([]string{"LiveDetectUrlReq", string(data)}, " ")
}
