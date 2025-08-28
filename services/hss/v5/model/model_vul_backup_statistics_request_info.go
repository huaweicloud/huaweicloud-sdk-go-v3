package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulBackupStatisticsRequestInfo 漏洞操作列表
type VulBackupStatisticsRequestInfo struct {

	// **参数解释**: 选择全部漏洞类型，该参数优先级高于data_list和host_data_list **约束限制**: 不涉及 **取值范围**: - all_vul：选择全部漏洞 - all_host：选择全部主机漏洞  **默认取值**: 不涉及
	SelectType *string `json:"select_type,omitempty"`

	// **参数解释**: 漏洞类型，当select_type选择all_vul时，需要传递该参数 **约束限制**: 不涉及 **取值范围**: - linux_vul : 漏洞类型-linux漏洞 - windows_vul : 漏洞类型-windows漏洞 - web_cms : Web-CMS漏洞 - app_vul : 应用漏洞 - urgent_vul : 应急漏洞  **默认取值**: linux_vul
	Type *string `json:"type,omitempty"`

	// **参数解释**: 漏洞列表 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	DataList *[]VulOperateInfo `json:"data_list,omitempty"`

	// **参数解释**: 主机维度漏洞列表 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	HostDataList *[]HostVulOperateInfo `json:"host_data_list,omitempty"`
}

func (o VulBackupStatisticsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulBackupStatisticsRequestInfo struct{}"
	}

	return strings.Join([]string{"VulBackupStatisticsRequestInfo", string(data)}, " ")
}
