package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteBackupRequestBody struct {

	// 需要批量删除的手动备份id列表，最多一次不超过10条。
	BackupIds []string `json:"backup_ids"`
}

func (o BatchDeleteBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBackupRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteBackupRequestBody", string(data)}, " ")
}
