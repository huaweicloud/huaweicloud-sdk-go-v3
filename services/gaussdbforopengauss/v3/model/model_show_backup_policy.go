package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupPolicy 备份策略信息。
type ShowBackupPolicy struct {

	// 全量备份文件可以保存的天数。
	KeepDays int32 `json:"keep_days"`

	// 全量备份时间段。自动备份将在该时间段内触发。除关闭自动备份策略外，必选。  取值范围：格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。  - HH取值必须比hh大1。 - mm和MM取值必须相同，且取值必须为00。
	StartTime string `json:"start_time"`

	// 全量备份周期配置。自动备份将在每星期指定的天进行。  取值范围：格式为逗号隔开的数字，数字代表星期。
	Period string `json:"period"`

	// 差量备份周期配置。自动差量备份将每隔周期分钟执行(废弃)。
	DifferentialPriod *string `json:"differential_priod,omitempty"`

	// 差量备份周期配置。自动差量备份将每隔周期分钟执行。
	DifferentialPeriod int32 `json:"differential_period"`

	// 备份时备份数据上传OBS的速度，单位为MB/s。范围为0~1024MB/s，默认75MB/s，0MB/s表示不限速。
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// 控制差量备份时读取磁盘上表文件差量修改页面的预取页面个数，可设置范围为1~8192，默认64。
	PrefetchBlock *int32 `json:"prefetch_block,omitempty"`

	// 废弃。
	FilesplitSize *int32 `json:"filesplit_size,omitempty"`

	// 全量、差量备份时产生的备份文件会根据分片大小进行拆分，可设置范围为0~1024GB，设置需为4的倍数，默认4GB，0GB表示不限制大小。  取值范围：0 ~ 1024
	FileSplitSize *int32 `json:"file_split_size,omitempty"`

	// 是否启用备机备份。  取值范围：true|false
	EnableStandbyBackup *bool `json:"enable_standby_backup,omitempty"`
}

func (o ShowBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicy struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicy", string(data)}, " ")
}
