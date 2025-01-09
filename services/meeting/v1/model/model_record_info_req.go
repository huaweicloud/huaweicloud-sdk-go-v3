package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordInfoReq 录制会议文件信息请求体
type RecordInfoReq struct {

	// 会议uuid
	ConfUUID string `json:"confUUID"`

	// 录制段落查询偏移量
	SegmentOffset *int32 `json:"segmentOffset,omitempty"`

	// 录制段落查询数量
	SegmentLimit *int32 `json:"segmentLimit,omitempty"`
}

func (o RecordInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordInfoReq struct{}"
	}

	return strings.Join([]string{"RecordInfoReq", string(data)}, " ")
}
