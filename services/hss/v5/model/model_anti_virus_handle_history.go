package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusHandleHistory 病毒查杀历史处置记录
type AntiVirusHandleHistory struct {

	// 病毒查杀结果ID
	ResultId *string `json:"result_id,omitempty"`

	// 病毒类型
	MalwareType *string `json:"malware_type,omitempty"`

	// 病毒名称
	MalwareName *string `json:"malware_name,omitempty"`

	// 威胁等级，包含如下:   - Security：安全   - Low: 低危   - Medium: 中危   - High: 高危   - Critical: 危急
	Severity *string `json:"severity,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 发生时间，毫秒
	OccurTime *int64 `json:"occur_time,omitempty"`

	// 处置状态，包含如下:   - unhandled：未处理   - handled: 已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// 处理方式，包含如下:   - mark_as_handled : 手动处理   - ignore : 忽略   - add_to_alarm_whitelist : 加入告警白名单   - isolate_and_kill : 隔离文件
	HandleMethod *string `json:"handle_method,omitempty"`

	// 备注信息
	Notes *string `json:"notes,omitempty"`

	// 处置时间
	HandleTime *int64 `json:"handle_time,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o AntiVirusHandleHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusHandleHistory struct{}"
	}

	return strings.Join([]string{"AntiVirusHandleHistory", string(data)}, " ")
}
