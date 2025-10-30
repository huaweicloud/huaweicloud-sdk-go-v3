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

	// 是否开启动态诱饵，包含如下2种，默认为关闭防护动态诱饵防护。   - opened ：开启。   - closed ：关闭。
	DeployMode *string `json:"deploy_mode,omitempty"`

	// 防护目录，多个目录请用英文分号隔开，最多支持填写20个防护目录
	ProtectionDirectory string `json:"protection_directory"`

	// **参数解释**: 需要防护的文件类型 **约束限制**: 不涉及 **取值范围**: txt、csv、rtf、doc、docx、xls、xlsx、ppt、pptx、pdf、xml、json、sql、mdf、dbf、ldf、db、myd、wdb、si、cfs、cfe、fnm、fdx、fdt、tvx、tvf、tvd、tim、nvd、nvm、dvd、dvm、jpeg、bmp、gif、png、tiff、eps、mp3、mp4、avi、mpg、wmv、RMVB、mov、3pg、swf、flv、rar、gz、tgz、zip、7z、cpp、c、java、asp、php、python、html、js、vdi、vmdk、vdx、ovf、qcow2、vmem、vswp、img、bak、back、cer、crt、pem、key、csr **默认取值**: 不涉及
	ProtectionType string `json:"protection_type"`

	// 排除目录(选填)，多个目录请用英文分号隔开，最多支持填写20个排除目录
	ExcludeDirectory *string `json:"exclude_directory,omitempty"`

	// 操作系统，包含如下：   - Windows : Windows系统   - Linux : Linux系统
	OperatingSystem string `json:"operating_system"`

	// 进程白名单
	ProcessWhitelist *[]TrustProcessInfo `json:"process_whitelist,omitempty"`

	// **参数解释**: 是否开启AI勒索防护 **约束限制**: 不涉及 **取值范围**: 包含两种：   - opened ：开启。   - closed ：关闭。 **默认取值**: closed
	AiProtectionStatus *string `json:"ai_protection_status,omitempty"`
}

func (o AddProtectionPolicyInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectionPolicyInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"AddProtectionPolicyInfoRequestInfo", string(data)}, " ")
}
