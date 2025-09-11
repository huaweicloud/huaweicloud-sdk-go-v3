package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditAutoBackupTemplateResponse Response Object
type SetAuditAutoBackupTemplateResponse struct {

	// OBS桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// OBS备份根路径
	BucketRootPath *string `json:"bucket_root_path,omitempty"`

	// 周期 - PER_DAY：每天 - PER_WEEK：每周 - PER_MONTH：每月 - PER_HOUR：每小时 - FIVE_MIN：每5分钟
	Cycle *string `json:"cycle,omitempty"`

	// 最近备份时间
	LatestBackupTime *sdktime.SdkTime `json:"latest_backup_time,omitempty"`

	// 备份开关  - 0：关闭  - 1：开启
	Status *int32 `json:"status,omitempty"`

	// 触发时间
	TriggerTime *sdktime.SdkTime `json:"trigger_time,omitempty"`

	// 是否已触发
	Triggered      *bool `json:"triggered,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SetAuditAutoBackupTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditAutoBackupTemplateResponse struct{}"
	}

	return strings.Join([]string{"SetAuditAutoBackupTemplateResponse", string(data)}, " ")
}
