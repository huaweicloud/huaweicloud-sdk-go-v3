package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InferenceVideoMarkInfo 分身数字人推理预处理视频标记信息。
type InferenceVideoMarkInfo struct {

	// 选取推理数据预处理视频起始时间。格式：“HH:MM:SS.mmm”。
	VideoStartTime *string `json:"video_start_time,omitempty"`

	// 选取推理数据预处理视频结束时间。格式：“HH:MM:SS.mmm”。
	VideoEndTime *string `json:"video_end_time,omitempty"`
}

func (o InferenceVideoMarkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InferenceVideoMarkInfo struct{}"
	}

	return strings.Join([]string{"InferenceVideoMarkInfo", string(data)}, " ")
}
