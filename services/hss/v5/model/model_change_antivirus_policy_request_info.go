package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAntivirusPolicyRequestInfo 修改自定义扫描策略
type ChangeAntivirusPolicyRequestInfo struct {

	// 策略ID
	PolicyId string `json:"policy_id"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	// **参数解释**： 启动类型 **取值范围**： 包含如下   - now：立即启动   - later：稍后启动   - period：周期启动
	StartType string `json:"start_type"`

	// **参数解释**： 任务类型 **取值范围**： 包含如下:   - quick ：快速扫描   - full : 全盘扫描   - custom : 自定义扫描
	ScanType *string `json:"scan_type,omitempty"`

	// 启动类型，包含如下:   - day ：每天   - week : 每周   - month : 每月
	ScanPeriod *string `json:"scan_period,omitempty"`

	// 扫描周期日期（1-28；扫描周期为week时，1-7代表周日至周六；扫描周期为month时，1-28代表每月1日到28日）
	ScanPeriodDate *int32 `json:"scan_period_date,omitempty"`

	// 扫描时间戳，毫秒（仅启动类型为later时有值）
	ScanTime *int64 `json:"scan_time,omitempty"`

	// 扫描小时数（仅启动类型为period时有值）
	ScanHour *int32 `json:"scan_hour,omitempty"`

	// 扫描分钟数（仅启动类型为period时有值）
	ScanMinute *int32 `json:"scan_minute,omitempty"`

	// 时区偏移量（仅启动类型为period时有值，单位：分钟）
	TimezoneOffset *int32 `json:"timezone_offset,omitempty"`

	// 文件类型集合型，包含如下:   - 0 ：全部   - 1 : 可执行   - 2 : 压缩   - 3 : 脚本   - 4 : 文档   - 5 : 图片   - 6 : 音视频
	FileTypes []int32 `json:"file_types"`

	// 扫描目录，多个用;分隔
	ScanDir *string `json:"scan_dir,omitempty"`

	// 排除目录，多个用;分隔
	IgnoreDir *string `json:"ignore_dir,omitempty"`

	// **参数解释**: 处置动作 **取值范围**: - auto：自动处置 - manual：人工处置
	Action string `json:"action"`

	// 此次扫描任务是否消耗按次计费配额
	WhetherPaidTask bool `json:"whether_paid_task"`

	// 策略管理主机列表
	HostIds []string `json:"host_ids"`
}

func (o ChangeAntivirusPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAntivirusPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeAntivirusPolicyRequestInfo", string(data)}, " ")
}
