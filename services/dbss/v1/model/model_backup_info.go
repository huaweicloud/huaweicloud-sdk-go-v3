package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupInfo struct {

	// 备份失败原因
	BackupTaskFailReason *string `json:"backup_task_fail_reason,omitempty"`

	// 备份时间,yyyy-MM-dd HH:mm:ss
	BackupTime *sdktime.SdkTime `json:"backup_time,omitempty"`

	// 标记删除
	Deleted *bool `json:"deleted,omitempty"`

	// 备份结束时间,yyyy-MM-dd HH:mm:ss
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 文件大小
	FileSize *int64 `json:"file_size,omitempty"`

	// 文件大小单位 - Byte - KB - MB - GB
	FileSizeUnit *string `json:"file_size_unit,omitempty"`

	// 备份ID
	Id *string `json:"id,omitempty"`

	// 备份方式 - AUTO：自动备份
	Mode *string `json:"mode,omitempty"`

	// 备份名称
	Name *string `json:"name,omitempty"`

	// 备份进度
	Percentage *int32 `json:"percentage,omitempty"`

	// 进度 - CLEAN_AFTER_FAILED: 清理失败 - DELETING：删除中 - DELETED：已删除 - DELETE_FAIL：删除失败 - RESTORING_WAITING：恢复等待中 - RESTORING：恢复中 - RESTORED：已恢复 - RESTORE_FAIL：恢复失败 - BACKUP_WAITING：等待备份 - FILE_UPLOAD_WAITING：等待上传备份文件 - FILE_UPLOADING：上传中 - AUTO_BACKUPING：自动备份中 - AUTO_BACKUPED：自动备份结束 - AUTO_BACKUP_FAIL：自动备份失败 - MANUAL_BACKUPING：手动备份中 - MANUAL_BACKUPED：手动备份结束 - MANUAL_BACKUP_FAIL：手动备份失败 - ISAP_WAITING：ISAP等待备份 - ISAP_BACKUPING：ISAP备份中 - ISAP_BACKUPED：ISAP备份成功 - ISAP_BACKUP_FAIL：ISAP备份失败 - ISAP_FILE_UPLOAD_WAITING：ISAP等待上传备份文件 - ISAP_FILE_UPLOADING：ISAP上传中
	Progress *string `json:"progress,omitempty"`

	// 还原失败原因
	RestoreTaskFailReason *string `json:"restore_task_fail_reason,omitempty"`

	// 文件SHA256
	Sha256 *string `json:"sha256,omitempty"`

	// 备份开始时间,yyyy-MM-dd HH:mm:ss
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`
}

func (o BackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInfo struct{}"
	}

	return strings.Join([]string{"BackupInfo", string(data)}, " ")
}
