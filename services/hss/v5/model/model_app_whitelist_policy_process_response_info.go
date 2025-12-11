package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppWhitelistPolicyProcessResponseInfo 策略识别进程信息
type AppWhitelistPolicyProcessResponseInfo struct {

	// **参数解释**： 进程名称 **取值范围**： 字符长度1-128位
	ProcessName *string `json:"process_name,omitempty"`

	// **参数解释**： 进程路径 **取值范围**： 字符长度1-256位
	ProcessPath *string `json:"process_path,omitempty"`

	// 进程hash
	ProcessHash *string `json:"process_hash,omitempty"`

	// **参数解释**： 处理方式 **取值范围**: - confirmed：已确认 - unconfirmed：未确认
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**： 指定目录 **取值范围**： 字符长度1-512位
	SpecifiedPath *string `json:"specified_path,omitempty"`

	// **参数解释**： 进程命令行 **约束限制**： 不涉及
	Cmdline *string `json:"cmdline,omitempty"`

	// **参数解释**: 文件大小 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	FileSize *int64 `json:"file_size,omitempty"`

	// **参数解释**： 文件签名 **取值范围**： 字符长度1-128位
	FileSigner *string `json:"file_signer,omitempty"`

	// **参数解释**： 进程类型 **约束限制**: 不涉及 **取值范围**: - 1：系统程序 - 2：解释类程序 - 3：普通可执行程序
	ProcessType *int32 `json:"process_type,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 应用类型 **约束限制**： 不涉及
	AppType *string `json:"app_type,omitempty"`

	// **参数解释**: 白名单确认次数 **约束限制**: 不涉及
	WhitelistCount *int32 `json:"whitelist_count,omitempty"`

	// **参数解释**: 黑名单确认次数 **约束限制**: 不涉及
	BlacklistCount *int32 `json:"blacklist_count,omitempty"`

	// **参数解释**： 进程可信状态 **约束限制**: 不涉及 **取值范围**: - 0：情报 - 1：恶意软件 - 2：人工确认 - 3：自学习  **默认取值**: 不涉及
	TrustStatusSource *int32 `json:"trust_status_source,omitempty"`

	// **参数解释**： 进程可信状态 **约束限制**: 不涉及 **取值范围**: - trust：可信 - suspicious：可疑 - malicious：恶意 - unknown：未知  **默认取值**: 不涉及
	ProcessStatus *string `json:"process_status,omitempty"`
}

func (o AppWhitelistPolicyProcessResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppWhitelistPolicyProcessResponseInfo struct{}"
	}

	return strings.Join([]string{"AppWhitelistPolicyProcessResponseInfo", string(data)}, " ")
}
