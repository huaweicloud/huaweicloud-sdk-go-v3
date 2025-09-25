package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupPolicyInfo 备份策略信息
type BackupPolicyInfo struct {

	// **参数解释**: 备份文件可以保存的天数。 **约束限制**: 不涉及。 **取值范围**: 1-732天。如果需要延长保留时间请联系客服人员申请，自动备份最长可以保留36500天。 **默认取值**: 不涉及。
	KeepDays int32 `json:"keep_days"`

	// **参数解释**: 备份时间段。自动备份将在该时间段内触发。 **约束限制**: - 该时间段为执行备份的UTC时间段。比如备份时间是当地时间05:00-06:00，时区为UTC+08:00，则接口需要传入21:00-22:00；如果时区为UTC+04:00，接口传入01:00-02:00。 - 当填写start_time时，建议同时填写period，再下发修改，以免修改后的备份时间段和周期在您的时区不符合预期。  **取值范围**: 非空，格式必须为hh:mm-HH:MM且有效，当前时间指UTC时间。 HH取值必须比hh大1，mm和MM取值必须相同，且取值必须为00。 取值示例： 21:00-22:00 **默认取值**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 全量备份周期配置。自动全量备份将在每周对应的UTC日期进行。 **约束限制**: - 该时间段为执行备份的UTC日期。比如备份时间为当地时间周一、周二05:00-06:00，时区为UTC+08:00，则period传入1,7；如果时区为UTC+04:00，period传入1,2。 - 当填写period时，建议同时填写start_time，再下发修改，以免修改后的备份时间段和周期在您的时区不符合预期。  **取值范围**: 格式为逗号隔开的数字，数字代表星期，取1~7。 取值示例： - 1,2,3,4 表示备份周期配置为星期一、星期二、星期三和星期四。 - 1,2,3,4,5,6,7 则表示星期一至星期日每天执行一次自动备份。 - 1,3,5表示周一、周三、周五执行一次自动备份。  **默认取值**: 不涉及。
	Period string `json:"period"`

	// **参数解释**: 差异备份间隔时间配置。每次自动差异备份的间隔时间。 **约束限制**: 不涉及。 **取值范围**: 格式为逗号隔开的数字，数字代表星期，取1~7。 取值示例： 15、30、60、180、360、720、1440。单位：分钟。 **默认取值**: 不涉及。
	DifferentialPeriod string `json:"differential_period"`

	// **参数解释**: 备份限速。控制备份时备份数据上传OBS的速度，限速用于限制上传备份对上传带宽的影响。 **约束限制**: 不涉及。 **取值范围**: 0~1024MB/s，0表示不限速。 **默认取值**: 75MB/s
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// **参数解释**: 差量预取页面个数。控制差量备份时读取磁盘上表文件差量修改页面的预取页面个数。当差量修改页面非常集中时（如数据导入场景），可以适当调大该值；当差量修改页面非常分散时（如随机更新），可以适当调小该值。 **约束限制**: 不涉及。 **取值范围**: 1~8192 **默认取值**: 64
	PrefetchBlock *int32 `json:"prefetch_block,omitempty"`

	// **参数解释**: 文件拆分大小。全量、差量备份时产生的备份文件会根据该参数的值进行拆分。 **约束限制**: 需为4的倍数。 **取值范围**: 可设置范围为0~1024GB。0GB表示不限制大小。 **默认取值**: 4GB
	FileSplitSize *int32 `json:"file_split_size,omitempty"`

	// **参数解释**: 是否启用备机备份。 **约束限制**: 不支持单节点实例及V2.0-3.100.0以下的实例。 **取值范围**: - true：启用备机备份。 - false：不启用备机备份。  **默认取值**: 不涉及。
	EnableStandbyBackup *bool `json:"enable_standby_backup,omitempty"`
}

func (o BackupPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicyInfo struct{}"
	}

	return strings.Join([]string{"BackupPolicyInfo", string(data)}, " ")
}
