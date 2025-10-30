package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulAffectedStaticsRequest Request Object
type ShowVulAffectedStaticsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 选择全部漏洞类型 **约束限制**: 不涉及 **取值范围**: - all_vul：选择全部漏洞 - all_host：选择全部主机漏洞  **默认取值**: 不涉及
	SelectType *string `json:"select_type,omitempty"`

	// **参数解释**: 漏洞类型，当前select_type为all_vul此字段为必选 **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞 - windows_vul：windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞 - urgent_vul：应急漏洞 - cluster_vul：集群漏洞  **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**: 集群id **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 容器id集合 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值1000 **默认取值**: 不涉及
	ContainerIds *[]string `json:"container_ids,omitempty"`

	// **参数解释**: 是否是容器场景 **约束限制**: 不涉及 **取值范围**: - true：是容器场景 - false：不是容器场景 **默认取值**: false
	IsContainer *bool `json:"is_container,omitempty"`

	// **参数解释**: 漏洞id集合 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值1000 **默认取值**: 不涉及
	VulIds *[]string `json:"vul_ids,omitempty"`

	// **参数解释**: 主机id集合 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值1000 **默认取值**: 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**: 类别，默认为host **约束限制**: 不涉及 **取值范围**: - host：主机 - container：容器 - serverless：serverless场景  **默认取值**: host
	Category *string `json:"category,omitempty"`
}

func (o ShowVulAffectedStaticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulAffectedStaticsRequest struct{}"
	}

	return strings.Join([]string{"ShowVulAffectedStaticsRequest", string(data)}, " ")
}
