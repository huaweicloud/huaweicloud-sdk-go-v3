package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogArchiveDto 慢日志Top信息
type SlowLogArchiveDto struct {

	// ID
	Id *int64 `json:"id,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 日志开始时间
	LogStartTime *int64 `json:"log_start_time,omitempty"`

	// 日志结束时间
	LogEndTime *int64 `json:"log_end_time,omitempty"`

	// 文件大小
	FileSize *int64 `json:"file_size,omitempty"`
}

func (o SlowLogArchiveDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogArchiveDto struct{}"
	}

	return strings.Join([]string{"SlowLogArchiveDto", string(data)}, " ")
}
