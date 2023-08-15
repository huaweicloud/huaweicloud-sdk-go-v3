package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoModerationResultRequestParamsData 创建作业时传的data参数
type VideoModerationResultRequestParamsData struct {

	// 创建作业时传的url参数
	Url string `json:"url"`

	// 创建作业时传的frame_interval参数，默认为5秒截取一帧
	FrameInterval *int32 `json:"frame_interval,omitempty"`
}

func (o VideoModerationResultRequestParamsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationResultRequestParamsData struct{}"
	}

	return strings.Join([]string{"VideoModerationResultRequestParamsData", string(data)}, " ")
}
