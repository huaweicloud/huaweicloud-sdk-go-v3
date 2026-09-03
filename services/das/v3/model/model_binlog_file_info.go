package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BinlogFileInfo binlog文件信息
type BinlogFileInfo struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 备份ID
	BackupId *string `json:"backup_id,omitempty"`

	// 文件大小
	FileSize *int64 `json:"file_size,omitempty"`

	TaskInfo *BinlogParseTaskInfo `json:"task_info,omitempty"`

	// binlog备份开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// binlog备份结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o BinlogFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BinlogFileInfo struct{}"
	}

	return strings.Join([]string{"BinlogFileInfo", string(data)}, " ")
}
