package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BinlogParseTaskInfo binlog解析任务信息
type BinlogParseTaskInfo struct {

	// 任务ID
	Id *int64 `json:"id,omitempty"`

	// 任务创建时间，单位：毫秒
	GmtCreate *int64 `json:"gmt_create,omitempty"`

	// 任务修改时间，单位：毫秒
	GmtModified *int64 `json:"gmt_modified,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 连接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// binlog类型。取值范围：latest（最近日志）、backup（归档日志）
	BinlogType *string `json:"binlog_type,omitempty"`

	// binlog文件名称
	FileName *string `json:"file_name,omitempty"`

	// 备份文件ID
	BackupId *string `json:"backup_id,omitempty"`

	// 任务状态。取值范围：0（初始化）、1（运行中）、2（部分成功）、3（成功）、4（失败）、-1（已删除）
	Status *int32 `json:"status,omitempty"`

	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty"`
}

func (o BinlogParseTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BinlogParseTaskInfo struct{}"
	}

	return strings.Join([]string{"BinlogParseTaskInfo", string(data)}, " ")
}
