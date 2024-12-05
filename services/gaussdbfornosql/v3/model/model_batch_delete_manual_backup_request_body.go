package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteManualBackupRequestBody struct {

	// 需要批量删除的手动备份id列表，最多一次不超过10条
	BackupIds []string `json:"backup_ids"`
}

func (o BatchDeleteManualBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteManualBackupRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteManualBackupRequestBody", string(data)}, " ")
}
