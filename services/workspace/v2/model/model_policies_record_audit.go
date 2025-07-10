package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesRecordAudit 录屏审计。
type PoliciesRecordAudit struct {

	// 是否开启录屏审计。取值为： false：表示关闭。 true：表示开启。
	Enable *bool `json:"enable,omitempty"`

	Rules *PoliciesRecordAuditRules `json:"rules,omitempty"`

	// 存储类型。取值为： OBS：OBS桶类型。 SFTP：SFTP对接类型。
	StorageType *string `json:"storage_type,omitempty"`

	// OBS桶来源。取值为： AUTO_CREATE：自动创建。 CREATED：已创建的。
	ObsBucketSource *string `json:"obs_bucket_source,omitempty"`

	// obs 桶名
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// 录屏文件保留时长（天）。取值为1~180天，0 表示永久保留。
	RetentionDuration *int32 `json:"retention_duration,omitempty"`
}

func (o PoliciesRecordAudit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesRecordAudit struct{}"
	}

	return strings.Join([]string{"PoliciesRecordAudit", string(data)}, " ")
}
