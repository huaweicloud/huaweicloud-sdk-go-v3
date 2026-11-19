package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RaspPatchInfo 虚拟补丁信息
type RaspPatchInfo struct {

	// **参数解释**: 特征检查项ID **取值范围**: 最小值0，最大值2097152
	ChkFeatureId *int64 `json:"chk_feature_id,omitempty"`

	// **参数解释**: 补丁名称 **取值范围**: 字符长度0-128位
	PatchName *string `json:"patch_name,omitempty"`

	// **参数解释**: 漏洞ID **取值范围**: 字符长度0-64位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞名称 **取值范围**: 字符长度0-512位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: CVE列表 **取值范围**: 0-128个RaspCveInfo对象
	CveList *[]RaspCveInfo `json:"cve_list,omitempty"`

	// **参数解释**: 目标组件 **取值范围**: 字符长度0-128位
	Artifact *string `json:"artifact,omitempty"`

	// **参数解释**: 影响版本 **取值范围**: 字符长度0-128位
	VulnVersion *string `json:"vuln_version,omitempty"`

	// **参数解释**: 补丁版本 **取值范围**: 字符长度0-64位
	PluginVersion *string `json:"plugin_version,omitempty"`

	// **参数解释**: 漏洞说明 **取值范围**: 字符长度0-512位
	Description *string `json:"description,omitempty"`

	// **参数解释**: 可选防护动作 **取值范围**: 最小值1，最大值3
	OptionalProtectiveAction *int32 `json:"optional_protective_action,omitempty"`

	// **参数解释**: 防护动作 **取值范围**: 最小值1，最大值2
	ProtectiveAction *int32 `json:"protective_action,omitempty"`
}

func (o RaspPatchInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RaspPatchInfo struct{}"
	}

	return strings.Join([]string{"RaspPatchInfo", string(data)}, " ")
}
