package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportImportFileResult 对象导入信息。
type SupportImportFileResult struct {

	// 文件导入阈值。
	FileSize *string `json:"file_size,omitempty"`

	// 上一次选择对象的方式。
	PreviousSelect *string `json:"previous_select,omitempty"`
}

func (o SupportImportFileResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportImportFileResult struct{}"
	}

	return strings.Join([]string{"SupportImportFileResult", string(data)}, " ")
}
