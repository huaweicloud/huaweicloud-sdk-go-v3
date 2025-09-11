package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupSwitchRequest struct {

	// OBS桶名称
	BucketName string `json:"bucket_name"`

	// 备份根路径
	BucketRootPath *string `json:"bucket_root_path,omitempty"`

	// 备份周期 - PER_DAY：每天 - PER_WEEK：每周 - PER_MONTH：每月 - PER_HOUR：每小时 - FIVE_MIN：每5分钟
	Cycle string `json:"cycle"`

	// 备份模式  - AUTO：自动备份
	Mode string `json:"mode"`

	// 开关状态  - 0：关闭  - 1：开启
	Status int32 `json:"status"`

	// 触发时间，yyyy-MM-dd HH:mm:ss
	TriggerTime *sdktime.SdkTime `json:"trigger_time"`
}

func (o BackupSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupSwitchRequest struct{}"
	}

	return strings.Join([]string{"BackupSwitchRequest", string(data)}, " ")
}
