package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RiskBackupConfig struct {

	// 账户ID
	BackupDomainId *string `json:"backup_domain_id,omitempty"`

	// OBS桶名称
	BucketName string `json:"bucket_name"`

	// OBS桶根路径
	BucketRootPath *string `json:"bucket_root_path,omitempty"`
}

func (o RiskBackupConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskBackupConfig struct{}"
	}

	return strings.Join([]string{"RiskBackupConfig", string(data)}, " ")
}
