package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbPolicyDetail struct {
	BackupPolicy *DbPolicyDetailBackupPolicy `json:"backup_policy,omitempty"`

	// 数据库复制策略
	OffsiteBackupPolicy *[]DbPolicyDetailOffsiteBackupPolicy `json:"offsite_backup_policy,omitempty"`
}

func (o DbPolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbPolicyDetail struct{}"
	}

	return strings.Join([]string{"DbPolicyDetail", string(data)}, " ")
}
