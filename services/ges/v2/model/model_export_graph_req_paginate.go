package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportGraphReqPaginate 分页相关参数，自2.3.11版本起导出图默认分页导出。
type ExportGraphReqPaginate struct {

	// 是否开启分页，默认为true，不需要开启分页时，需显示声明为false。
	Enable *bool `json:"enable,omitempty"`

	// 按页导出时，每个文件最大行数，默认10000000。
	RowCountPerFile *int32 `json:"row_count_per_file,omitempty"`

	// 按页导出时，并行线程数，默认为8。
	NumThread *int32 `json:"num_thread,omitempty"`
}

func (o ExportGraphReqPaginate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportGraphReqPaginate struct{}"
	}

	return strings.Join([]string{"ExportGraphReqPaginate", string(data)}, " ")
}
