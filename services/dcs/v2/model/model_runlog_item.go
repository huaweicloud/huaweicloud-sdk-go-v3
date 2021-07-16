package model

import (
	"encoding/json"

	"strings"
)

// 运行日志单个条目
type RunlogItem struct {
	// 日志的唯一标识

	Id *string `json:"id,omitempty"`
	// 运行日志文件名

	FileName *string `json:"file_name,omitempty"`
	// 获取运行日志状态

	Status *string `json:"status,omitempty"`
	// 运行日志采集的日期，格式为\"yyyy-MM-dd\"

	Time *string `json:"time,omitempty"`
}

func (o RunlogItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunlogItem struct{}"
	}

	return strings.Join([]string{"RunlogItem", string(data)}, " ")
}
