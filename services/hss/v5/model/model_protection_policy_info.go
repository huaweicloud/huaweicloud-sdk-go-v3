package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectionPolicyInfo struct {

	// **参数解释**: 策略ID **取值范围**: 字符长度0-128
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**: 防护策略名称 **取值范围**: 字符长度1-128
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 防护动作 **取值范围**: 包含如下2种。   - alarm_and_isolation ：告警并自动隔离。   - alarm_only ：仅告警。
	ProtectionMode *string `json:"protection_mode,omitempty"`

	// **参数解释**: 是否开启诱饵防护 **取值范围**: 包含如下1种，默认为开启防护诱饵防护。   - opened ：开启。
	BaitProtectionStatus *string `json:"bait_protection_status,omitempty"`

	// **参数解释**: 是否开启动态诱饵防护 **取值范围**: 包含如下2种，默认为关闭动态诱饵防护。   - opened ：开启。   - closed ：关闭。
	DeployMode *string `json:"deploy_mode,omitempty"`

	// **参数解释**: 防护目录 **取值范围**: 字符长度1-128
	ProtectionDirectory *string `json:"protection_directory,omitempty"`

	// **参数解释**: 需要防护的文件类型 **取值范围**: txt、csv、rtf、doc、docx、xls、xlsx、ppt、pptx、pdf、xml、json、sql、mdf、dbf、ldf、db、myd、wdb、si、cfs、cfe、fnm、fdx、fdt、tvx、tvf、tvd、tim、nvd、nvm、dvd、dvm、jpeg、bmp、gif、png、tiff、eps、mp3、mp4、avi、mpg、wmv、RMVB、mov、3pg、swf、flv、rar、gz、tgz、zip、7z、cpp、c、java、asp、php、python、html、js、vdi、vmdk、vdx、ovf、qcow2、vmem、vswp、img、bak、back、cer、crt、pem、key、csr
	ProtectionType *string `json:"protection_type,omitempty"`

	// **参数解释**: 排除目录，选填 **取值范围**: 字符长度1-128
	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	// **参数解释**: 是否运行时检测 **取值范围**: 包含如下2种，暂时只有关闭一种状态，为保留字段。   - opened ：开启。   - closed ：关闭。
	RuntimeDetectionStatus *string `json:"runtime_detection_status,omitempty"`

	// **参数解释**: 运行时检测目录，现在为保留字段 **取值范围**: 字符长度1-128
	RuntimeDetectionDirectory *string `json:"runtime_detection_directory,omitempty"`

	// **参数解释**: 关联server个数 **取值范围**: 取值范围0-2097152
	CountAssociatedServer *int32 `json:"count_associated_server,omitempty"`

	// **参数解释**: 操作系统类型。 - Linux - Windows **取值范围**: 字符长度1-128
	OperatingSystem *string `json:"operating_system,omitempty"`

	// 进程白名单
	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`

	// **参数解释**: 是否为默认策略 **取值范围**: 包含如下2种。   - 0 ：非默认策略。   - 1 ：默认策略
	DefaultPolicy *int32 `json:"default_policy,omitempty"`

	// **参数解释**: 是否开启AI勒索防护 **取值范围**:   - opened ：开启。   - closed ：关闭。
	AiProtectionStatus *string `json:"ai_protection_status,omitempty"`
}

func (o ProtectionPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionPolicyInfo struct{}"
	}

	return strings.Join([]string{"ProtectionPolicyInfo", string(data)}, " ")
}
