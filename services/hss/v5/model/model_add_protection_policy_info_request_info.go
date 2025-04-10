package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddProtectionPolicyInfoRequestInfo struct {

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 防护动作，包含如下2种。   - alarm_and_isolation ：告警并自动隔离。   - alarm_only ：仅告警。
	ProtectionMode string `json:"protection_mode"`

	// 是否开启动态诱饵，包含如下2种, 默认为关闭防护动态诱饵防护。   - opened ：开启。   - closed ：关闭。
	DeployMode *string `json:"deploy_mode,omitempty"`

	// 防护目录,多个目录请用英文分号隔开，最多支持填写20个防护目录
	ProtectionDirectory string `json:"protection_directory"`

	// 防护文件类型，例如：docx，txt，avi
	ProtectionType string `json:"protection_type"`

	// 排除目录(选填)，多个目录请用英文分号隔开，最多支持填写20个排除目录
	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	// 操作系统，包含如下：   - Windows : Windows系统   - Linux : Linux系统
	OperatingSystem string `json:"operating_system"`

	// 进程白名单
	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`

	// 是否开启AI勒索防护，包含如下2种, 默认为关闭AI勒索防护，当前只支持Windows防护策略   - opened ：开启。   - closed ：关闭。
	AiProtectionStatus *string `json:"ai_protection_status,omitempty"`
}

func (o AddProtectionPolicyInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectionPolicyInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"AddProtectionPolicyInfoRequestInfo", string(data)}, " ")
}
