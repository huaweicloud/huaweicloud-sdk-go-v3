package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileOperateLog struct {
	// 传输时间 格式：hh:ii:ss

	Duration *string `json:"duration,omitempty"`
	// 操作时间

	OperateTime *string `json:"operate_time,omitempty"`
	// 操作类型

	OperateType *string `json:"operate_type,omitempty"`
	// 文件名

	FileName *string `json:"file_name,omitempty"`
	// 来源路径

	FromPath *string `json:"from_path,omitempty"`
	// 目标路径

	ToPath *string `json:"to_path,omitempty"`
	// 文件大小，多少k，多少M，多少G。

	FileSize *string `json:"file_size,omitempty"`
	// 操作结果

	OperResult *string `json:"oper_result,omitempty"`
}

func (o FileOperateLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileOperateLog struct{}"
	}

	return strings.Join([]string{"FileOperateLog", string(data)}, " ")
}
