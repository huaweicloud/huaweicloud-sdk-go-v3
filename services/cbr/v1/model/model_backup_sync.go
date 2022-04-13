package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupSync struct {
	// 备份副本ID

	BackupId string `json:"backup_id"`
	// 备份名称

	BackupName string `json:"backup_name"`
	// 桶名

	BucketName string `json:"bucket_name"`
	// 备份链在存储单元上的路径

	ImagePath string `json:"image_path"`
	// 备份对象ID

	ResourceId string `json:"resource_id"`
	// 备份对象名称

	ResourceName string `json:"resource_name"`
	// 备份对象资源类型

	ResourceType string `json:"resource_type"`
	// 备份时间戳，例如1548898428

	CreatedAt int32 `json:"created_at"`
}

func (o BackupSync) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupSync struct{}"
	}

	return strings.Join([]string{"BackupSync", string(data)}, " ")
}
