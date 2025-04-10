package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusPolicyResponseInfo 策略信息
type AntiVirusPolicyResponseInfo struct {

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 启动类型，包含如下:   - now : 立即启动   - later : 稍后启动   - period : 周期启动
	StartType *string `json:"start_type,omitempty"`

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

	// 下次启动时间，毫秒
	NextStartTime *int64 `json:"next_start_time,omitempty"`

	// 扫描目录，多个用;分隔
	ScanDir *string `json:"scan_dir,omitempty"`

	// 排除目录，多个用;分隔
	IgnoreDir *string `json:"ignore_dir,omitempty"`

	// 处置动作，包含如下:   - auto ：自动处置   - manual : 人工处置
	Action *string `json:"action,omitempty"`

	// 失效，包含如下:   - true ：是   - fasle ：否
	Invalidate *bool `json:"invalidate,omitempty"`

	// 关联服务器数
	HostNum *int32 `json:"host_num,omitempty"`

	// 主机信息
	HostInfoList *[]AntiVirusPolicyHostResponseInfo `json:"host_info_list,omitempty"`

	// 此次扫描任务是否付费
	WhetherPaidTask *bool `json:"whether_paid_task,omitempty"`

	// 文件类型集合型
	FileTypeList *[]int32 `json:"file_type_list,omitempty"`
}

func (o AntiVirusPolicyResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusPolicyResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusPolicyResponseInfo", string(data)}, " ")
}
