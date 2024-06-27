package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupPolicyInfo 备份策略信息
type BackupPolicyInfo struct {

	// 备份文件可以保存的天数。取值范围：1-36500天。
	KeepDays int32 `json:"keep_days"`

	// 备份时间段。自动备份将在该时间段内触发。取值范围：非空，格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。 HH取值必须比hh大1，mm和MM取值必须相同，且取值必须为00。 取值示例： 21:00-22:00
	StartTime string `json:"start_time"`

	// 全量备份周期配置。自动全量备份将在每星期指定的天进行。 取值范围：格式为逗号隔开的数字，数字代表星期。取值示例：1,2,3,4则表示备份周期配置为星期一、星期二、星期三和星期四。
	Period string `json:"period"`

	// 差异备份间隔时间配置。每次自动差异备份的间隔时间。 取值范围：15、30、60、180、360、720、1440。单位：分钟。 取值示例：30
	DifferentialPeriod string `json:"differential_period"`

	// 备份限速,默认值为0MB/s，表示不限速。控制备份是备份数据上传OBS的速度，限速用于限制上传备份对上传带宽的影响 取值范围：0~ 1024 最小值：0 MB/s
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// 控制差量备份时读取磁盘上表文件差量修改页面的预取页面个数。当差量修改页面非常集中时（如数据导入场景），可以适当调大            该值；当差量修改页面非常分散时（如随机更新），可以适当调小该值。默认为64,单位个数 取值范围：1 ~ 8192 最小值：1 最大值：8192
	PrefetchBlock *int32 `json:"prefetch_block,omitempty"`

	// 全量、差量备份时产生的备份文件会根据该参数的值进行拆分，可设置范围为0~1024GB，设置需为4的倍数，默认4GB，0GB表示不           限制大小。 取值范围：0 ~ 1024 最小值：0 最大值：1024
	FileSplitSize *int32 `json:"file_split_size,omitempty"`

	// 是否启用备机备份。(不支持单节点实例及3.100.0以下的实例)
	EnableStandbyBackup *bool `json:"enable_standby_backup,omitempty"`
}

func (o BackupPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicyInfo struct{}"
	}

	return strings.Join([]string{"BackupPolicyInfo", string(data)}, " ")
}
