package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoStreamCreateRequestData 视频流数据输入
type VideoStreamCreateRequestData struct {

	// 视频流url，支持rtmp、rtmps、hls、http、https等主流协议。
	Url string `json:"url"`

	// 截帧频率间隔，单位为秒，取值范围为1~60s；若不传递默认5s截帧一次
	FrameInterval *int32 `json:"frame_interval,omitempty"`
}

func (o VideoStreamCreateRequestData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoStreamCreateRequestData struct{}"
	}

	return strings.Join([]string{"VideoStreamCreateRequestData", string(data)}, " ")
}
