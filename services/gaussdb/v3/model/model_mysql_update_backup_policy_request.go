package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlUpdateBackupPolicyRequest 设置自动备份策略信息
type MysqlUpdateBackupPolicyRequest struct {
	BackupPolicy *MysqlBackupPolicy `json:"backup_policy"`
}

func (o MysqlUpdateBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlUpdateBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"MysqlUpdateBackupPolicyRequest", string(data)}, " ")
}
