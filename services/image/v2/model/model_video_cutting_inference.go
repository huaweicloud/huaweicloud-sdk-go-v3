package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 视频剪切服务推理分段参数
type VideoCuttingInference struct {

	// 视频剪切服务推理分段参数
	SegmentInfo []VideoSegmentInfo `json:"segment_info"`
}

func (o VideoCuttingInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCuttingInference struct{}"
	}

	return strings.Join([]string{"VideoCuttingInference", string(data)}, " ")
}
