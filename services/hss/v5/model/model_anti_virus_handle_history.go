package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusHandleHistory 病毒查杀历史处置记录
type AntiVirusHandleHistory struct {

	// **参数解释**： 病毒查杀结果ID **取值范围**： 字符长度1-64位
	ResultId *string `json:"result_id,omitempty"`

	// **参数解释**： 病毒类型 **取值范围**： Trojan（木马）、Virus（病毒）、Worm（蠕虫）等
	MalwareType *string `json:"malware_type,omitempty"`

	// **参数解释**： 病毒名称 **取值范围**： 字符长度1-128位
	MalwareName *string `json:"malware_name,omitempty"`

	// **参数解释**： 威胁等级 **取值范围**： Security（安全）、Low（低危）、Medium（中危）、High（高危）、Critical（致命）
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位，支持IPv4或IPv6格式（IPv4长度7-15位，IPv6长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 资产重要性。 **取值范围**： - important ：重要资产。 - common ：一般资产。 - test ：测试资产。
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 发生时间，毫秒 **取值范围**： 最小值0，最大值9223372036854775807，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	OccurTime *int64 `json:"occur_time,omitempty"`

	// **参数解释**： 处理状态 **取值范围**： - unhandled：未处理 - handled：已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 处理方式 **取值范围**: 包含如下:   - mark_as_handled：手动处理   - ignore：忽略   - add_to_alarm_whitelist：加入告警白名单   - isolate_and_kill：隔离文件   - unhandle：取消手动处理   - do_not_ignore：取消忽略   - remove_from_alarm_whitelist：删除告警白名单   - do_not_isolate_or_kill：取消隔离文件
	HandleMethod *string `json:"handle_method,omitempty"`

	// **参数解释** 备注信息 **取值范围** 字符长度0-512位
	Notes *string `json:"notes,omitempty"`

	// **参数解释**: 处置时间 **取值范围**: 非负长整数，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算）；单位：ms
	HandleTime *int64 `json:"handle_time,omitempty"`

	// **参数解释**: 用户名 **取值范围**: 字符长度1-64位
	UserName *string `json:"user_name,omitempty"`
}

func (o AntiVirusHandleHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusHandleHistory struct{}"
	}

	return strings.Join([]string{"AntiVirusHandleHistory", string(data)}, " ")
}
