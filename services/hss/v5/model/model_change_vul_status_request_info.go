package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVulStatusRequestInfo 漏洞操作列表
type ChangeVulStatusRequestInfo struct {

	// **参数解释**: 对漏洞进行的处置操作类型 **约束限制**: 不涉及 **取值范围**: - ignore：忽略 - not_ignore：取消忽略 - immediate_repair：修复 - manual_repair：人工修复 - verify：验证 - add_to_whitelist：加入白名单  **默认取值**: 不涉及
	OperateType string `json:"operate_type"`

	// **参数解释**: 本次处置操作的备注信息 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位 **默认取值**: 不涉及
	Remark *string `json:"remark,omitempty"`

	// **参数解释**: 处置全部漏洞的类型 **约束限制**: 只有需要对全部漏洞进行处置时需要该参数 **取值范围**: - all_vul：按照指定漏洞类型处置全部漏洞 - all_host：处置全部主机的漏洞  **默认取值**: 不涉及
	SelectType *string `json:"select_type,omitempty"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：漏洞类型-linux漏洞 - windows_vul：漏洞类型-windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞 - urgent_vul：应急漏洞  **默认取值**: linux_vul
	Type *string `json:"type,omitempty"`

	// **参数解释**: 通过漏洞维度指定需要处置的漏洞信息 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值500 **默认取值**: 不涉及
	DataList *[]VulOperateInfo `json:"data_list,omitempty"`

	// **参数解释**: 通过主机维度指定需要处置的漏洞信息 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值500 **默认取值**: 不涉及
	HostDataList *[]HostVulOperateInfo `json:"host_data_list,omitempty"`

	// **参数解释**: 本次漏洞处置对应的备份信息id，若不传该参数，则不进行备份 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	BackupInfoId *string `json:"backup_info_id,omitempty"`

	// **参数解释**: 自定义备份主机使用的存储库及备份名称列表。不在该列表中的主机备份时系统会自动选取剩余空间最大的存储库，并自动生成备份名称 **约束限制**: 只有backup_info_id有值时该参数才会生效 **取值范围**: 最小值1，最大值50 **默认取值**: 不涉及
	CustomBackupHosts *[]ChangeVulStatusRequestInfoCustomBackupHosts `json:"custom_backup_hosts,omitempty"`
}

func (o ChangeVulStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeVulStatusRequestInfo", string(data)}, " ")
}
