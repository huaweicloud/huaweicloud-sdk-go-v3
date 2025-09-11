package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupRiskBucketPathResponse Response Object
type ShowBackupRiskBucketPathResponse struct {

	// 账户ID
	BackupDomainId *string `json:"backup_domain_id,omitempty"`

	// OBS桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// OBS桶根路径
	BucketRootPath *string `json:"bucket_root_path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackupRiskBucketPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRiskBucketPathResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupRiskBucketPathResponse", string(data)}, " ")
}
