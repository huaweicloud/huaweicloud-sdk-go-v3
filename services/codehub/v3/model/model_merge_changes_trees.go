package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeChangesTrees struct {

	// 分段路径
	Title *string `json:"title,omitempty"`

	// 路径级别
	Level *int32 `json:"level,omitempty"`

	// 当前级别全路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`
}

func (o MergeChangesTrees) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeChangesTrees struct{}"
	}

	return strings.Join([]string{"MergeChangesTrees", string(data)}, " ")
}
