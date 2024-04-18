package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceLogFile struct {

	// 日志采集状态。
	Status *string `json:"status,omitempty"`

	// 日志文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 日志开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 日志结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 日志文件大小，单位kb。
	FileSize *string `json:"file_size,omitempty"`

	// 日志文件下载链接。
	FileLink *string `json:"file_link,omitempty"`
}

func (o InstanceLogFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceLogFile struct{}"
	}

	return strings.Join([]string{"InstanceLogFile", string(data)}, " ")
}
