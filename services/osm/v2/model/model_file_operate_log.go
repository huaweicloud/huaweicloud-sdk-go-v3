package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileOperateLog struct {

	// 传输时间 格式：hh:ii:ss
	Duration *string `json:"duration,omitempty" xml:"duration"`

	// 操作时间
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time"`

	// 操作类型
	OperateType *string `json:"operate_type,omitempty" xml:"operate_type"`

	// 文件名
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 来源路径
	FromPath *string `json:"from_path,omitempty" xml:"from_path"`

	// 目标路径
	ToPath *string `json:"to_path,omitempty" xml:"to_path"`

	// 文件大小，多少k，多少M，多少G。
	FileSize *string `json:"file_size,omitempty" xml:"file_size"`

	// 操作结果
	OperResult *string `json:"oper_result,omitempty" xml:"oper_result"`
}

func (o FileOperateLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileOperateLog struct{}"
	}

	return strings.Join([]string{"FileOperateLog", string(data)}, " ")
}
