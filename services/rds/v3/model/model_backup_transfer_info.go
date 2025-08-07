package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupTransferInfo 转储任务信息
type BackupTransferInfo struct {

	// 备份id
	Id *string `json:"id,omitempty"`

	// 备份名称
	Name *string `json:"name,omitempty"`

	// 开始时间
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 对象大小,单位为KByte
	Size *float64 `json:"size,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 目标对象名称
	FileName *string `json:"file_name,omitempty"`

	// 目标桶名
	Destination *string `json:"destination,omitempty"`

	// 转储备份起始时间
	BackupBeginTime *int64 `json:"backup_begin_time,omitempty"`

	// 转储备份结束时间
	BackupEndTime *int64 `json:"backup_end_time,omitempty"`

	// 转储任务类型
	TransferType *string `json:"transfer_type,omitempty"`

	// 转储目标前缀
	Prefix *string `json:"prefix,omitempty"`

	// 转储备份类型
	Type *string `json:"type,omitempty"`
}

func (o BackupTransferInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTransferInfo struct{}"
	}

	return strings.Join([]string{"BackupTransferInfo", string(data)}, " ")
}
