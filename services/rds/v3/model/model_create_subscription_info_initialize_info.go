package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionInfoInitializeInfo 初始化信息。
type CreateSubscriptionInfoInitializeInfo struct {

	// 初始化使用的文件源，仅支持OBS或BACKUP。
	FileSource *string `json:"file_source,omitempty"`

	// 使用备份文件初始化的备份文件ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 使用OBS内备份文件恢复的bucket名称。
	BucketName *string `json:"bucket_name,omitempty"`

	// OBS桶内备份文件的路径。
	FilePath *string `json:"file_path,omitempty"`

	// OBS桶内备份文件的名称。
	FileName *string `json:"file_name,omitempty"`

	// 是否使用备份文件对订阅库进行覆盖还原。
	OverwriteRestore *bool `json:"overwrite_restore,omitempty"`
}

func (o CreateSubscriptionInfoInitializeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionInfoInitializeInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionInfoInitializeInfo", string(data)}, " ")
}
