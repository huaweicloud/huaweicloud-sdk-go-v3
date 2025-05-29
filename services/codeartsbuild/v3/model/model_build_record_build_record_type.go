package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildRecordBuildRecordType 构建记录类型
type BuildRecordBuildRecordType struct {

	// 是否rerun
	Rerun *bool `json:"rerun,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 记录类型
	RecordType *string `json:"record_type,omitempty"`

	// 是否返回
	IsRerun *bool `json:"is_rerun,omitempty"`
}

func (o BuildRecordBuildRecordType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildRecordBuildRecordType struct{}"
	}

	return strings.Join([]string{"BuildRecordBuildRecordType", string(data)}, " ")
}
