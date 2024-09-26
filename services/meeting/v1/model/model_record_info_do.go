package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordInfoDo 响应基础类
type RecordInfoDo struct {

	// 会议主题
	Subject *string `json:"subject,omitempty"`

	// 会议录制开始时间
	BeginTime *string `json:"beginTime,omitempty"`

	// 录制段落查询偏移量
	SegmentOffset *int32 `json:"segmentOffset,omitempty"`

	// 录制段落查询数量
	SegmentLimit *int32 `json:"segmentLimit,omitempty"`

	// 录制段落总数
	SegmentCount *int32 `json:"segmentCount,omitempty"`

	// 录制人工启动/停止分段列表
	SegmentList *[]SegmentDo `json:"segmentList,omitempty"`
}

func (o RecordInfoDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordInfoDo struct{}"
	}

	return strings.Join([]string{"RecordInfoDo", string(data)}, " ")
}
