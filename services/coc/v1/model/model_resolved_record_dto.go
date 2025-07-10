package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResolvedRecordDto 清除汇聚告警的记录
type ResolvedRecordDto struct {

	// 清除汇聚告警的时间
	ResolvedTime *int64 `json:"resolved_time,omitempty"`

	// 执行人
	CreateName *string `json:"create_name,omitempty"`

	// 清除时填写的备注
	Remarks *string `json:"remarks,omitempty"`
}

func (o ResolvedRecordDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolvedRecordDto struct{}"
	}

	return strings.Join([]string{"ResolvedRecordDto", string(data)}, " ")
}
