package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBinlogTaskRequestBody 创建binlog解析任务请求体
type CreateBinlogTaskRequestBody struct {

	// binlog类型。取值范围：latest（最近日志）、backup（归档日志）、fragment（碎片备份日志）
	BinlogType string `json:"binlog_type"`

	// binlog文件名称
	FileName string `json:"file_name"`

	// 归档ID
	BackupId *string `json:"backup_id,omitempty"`
}

func (o CreateBinlogTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBinlogTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBinlogTaskRequestBody", string(data)}, " ")
}
