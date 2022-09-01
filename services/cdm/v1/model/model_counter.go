package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Counter struct {

	// 写入的字节数
	BytesWritten *int64 `json:"BYTES_WRITTEN,omitempty" xml:"BYTES_WRITTEN"`

	// 总文件数
	TotalFiles *int32 `json:"TOTAL_FILES,omitempty" xml:"TOTAL_FILES"`

	// 读取的行数
	RowsRead *int64 `json:"ROWS_READ,omitempty" xml:"ROWS_READ"`

	// 读取的字节数
	BytesRead *int64 `json:"BYTES_READ,omitempty" xml:"BYTES_READ"`

	// 写入的行数
	RowsWritten *int64 `json:"ROWS_WRITTEN,omitempty" xml:"ROWS_WRITTEN"`

	// 写入的文件数
	FilesWritten *int32 `json:"FILES_WRITTEN,omitempty" xml:"FILES_WRITTEN"`

	// 读取的文件数
	FilesRead *int32 `json:"FILES_READ,omitempty" xml:"FILES_READ"`

	// 总字节数
	TotalSize *int64 `json:"TOTAL_SIZE,omitempty" xml:"TOTAL_SIZE"`

	// 跳过的文件数
	FilesSkipped *int32 `json:"FILES_SKIPPED,omitempty" xml:"FILES_SKIPPED"`

	// 跳过的行数
	RowsWrittenSkipped *int64 `json:"ROWS_WRITTEN_SKIPPED,omitempty" xml:"ROWS_WRITTEN_SKIPPED"`
}

func (o Counter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Counter struct{}"
	}

	return strings.Join([]string{"Counter", string(data)}, " ")
}
