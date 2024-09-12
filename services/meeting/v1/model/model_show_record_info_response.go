package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordInfoResponse Response Object
type ShowRecordInfoResponse struct {

	// 结果码
	Code *int32 `json:"code,omitempty"`

	// 结果描述
	Message *string `json:"message,omitempty"`

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
	SegmentList    *[]SegmentDo `json:"segmentList,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRecordInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordInfoResponse", string(data)}, " ")
}
