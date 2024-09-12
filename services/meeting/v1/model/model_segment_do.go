package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SegmentDo 录制切片段落
type SegmentDo struct {

	// 录制人工分段序号，每次启动录制，生成一个分段
	ManualOrder *int32 `json:"manualOrder,omitempty"`

	// 录制片段内的文件自动切片序号（每次启动录制后，每三小时一个分片文件，多次人工启动重新从1开始）
	SegmentOrder *int32 `json:"segmentOrder,omitempty"`

	// 录制分段总文件大小（字节）
	SegmentSize *string `json:"segmentSize,omitempty"`

	// 录制文件自动切片列表，每3小时文件切片一次
	FileList *[]SegmentFileDo `json:"fileList,omitempty"`
}

func (o SegmentDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SegmentDo struct{}"
	}

	return strings.Join([]string{"SegmentDo", string(data)}, " ")
}
