package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckRuleHostResponseInfo 受单个基线检查项影响的单个服务器信息
type SecurityCheckRuleHostResponseInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 配置检查（基线）的名称，例如SSH、CentOS 7、Windows **取值范围** 字符长度0-256位
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释** 基线的名称，例如SSH 应用基线检查、CentOS 7 系统基线检查、Windows 系统基线检查 **取值范围** 字符长度0-256位
	BaselineName *string `json:"baseline_name,omitempty"`

	// **参数解释**: 服务器公网IP **取值范围**: 字符长度0-128位
	HostPublicIp *string `json:"host_public_ip,omitempty"`

	// **参数解释** 服务器私有IP **取值范围** 字符长度0-256位
	HostPrivateIp *string `json:"host_private_ip,omitempty"`

	// **参数解释** 扫描时间(ms) **取值范围** 取值0-9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释** 风险项数量 **取值范围** 取值0-2147483647
	FailedNum *int32 `json:"failed_num,omitempty"`

	// **参数解释** 通过项数量 **取值范围** 取值0-2147483647
	PassedNum *int32 `json:"passed_num,omitempty"`

	// **参数解释** 差异化展示提示信息 **取值范围** 字符长度0-512位
	DiffDescription *string `json:"diff_description,omitempty"`

	// **参数解释** 忽略或加白的备注 **取值范围** 字符长度0-1024位
	Description *string `json:"description,omitempty"`

	// **参数解释** 主机类型 **取值范围** - cce
	HostType *string `json:"host_type,omitempty"`

	// **参数解释** 是否支持一键修复 **取值范围** - 1 : 支持一键修复 - 0 : 不支持
	EnableFix *int32 `json:"enable_fix,omitempty"`

	// **参数解释** 该检查项的修复&忽略&验证按钮是否可单击 **取值范围** - true  : 按钮可单击 - false : 按钮不可单击
	EnableClick *bool `json:"enable_click,omitempty"`

	// **参数解释** 已忽略检查项是否可点击 **取值范围** - true  : 按钮可单击 - false : 按钮不可单击
	CancelIgnoreEnableClick *bool `json:"cancel_ignore_enable_click,omitempty"`

	// **参数解释** 检测结果类型 **取值范围** - safe             : 已通过 - unhandled        : 未处理 - ignored          : 已忽略 - fixing           : 修复中 - fix-failed       : 修复失败 - verifying        : 验证中 - add_to_whitelist : 已加白(表示检测失败，但已进行加白)
	ResultType *string `json:"result_type,omitempty"`

	// **参数解释** 修复失败原因 **取值范围** 字符长度0-256位
	FixFailedReason *string `json:"fix_failed_reason,omitempty"`

	// **参数解释** 集群ID **取值范围** 字符长度0-64位
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o SecurityCheckRuleHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckRuleHostResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckRuleHostResponseInfo", string(data)}, " ")
}
