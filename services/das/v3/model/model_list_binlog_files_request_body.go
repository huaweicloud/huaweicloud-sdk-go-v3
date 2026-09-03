package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogFilesRequestBody 查询binlog文件列表请求体
type ListBinlogFilesRequestBody struct {

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 当前页
	CurPage *int32 `json:"cur_page,omitempty"`

	// 分页大小
	PerPage *int32 `json:"per_page,omitempty"`

	// binlog类型
	BinlogType string `json:"binlog_type"`
}

func (o ListBinlogFilesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogFilesRequestBody struct{}"
	}

	return strings.Join([]string{"ListBinlogFilesRequestBody", string(data)}, " ")
}
