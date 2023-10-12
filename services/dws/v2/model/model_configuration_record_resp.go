package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationRecordResp struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 操作
	Operator *string `json:"operator,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o ConfigurationRecordResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRecordResp struct{}"
	}

	return strings.Join([]string{"ConfigurationRecordResp", string(data)}, " ")
}
