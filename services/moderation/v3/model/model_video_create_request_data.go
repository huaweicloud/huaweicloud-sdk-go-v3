package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoCreateRequestData 视频数据输入
type VideoCreateRequestData struct {

	// 视频url地址
	Url string `json:"url"`

	// 截帧频率间隔，单位为秒，取值范围为1~60s；若不传递默认5s截帧一次
	FrameInterval *int32 `json:"frame_interval,omitempty"`
}

func (o VideoCreateRequestData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCreateRequestData struct{}"
	}

	return strings.Join([]string{"VideoCreateRequestData", string(data)}, " ")
}
