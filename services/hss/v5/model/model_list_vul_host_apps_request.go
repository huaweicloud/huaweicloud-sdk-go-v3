package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostAppsRequest Request Object
type ListVulHostAppsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 主机id **约束限制**: 不涉及 **取值范围**: 字符长度1-128 **默认取值**: 不涉及
	HostId string `json:"host_id"`

	// **参数解释**: 漏洞ID **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 漏洞处置状态 **约束限制**: 不涉及 **取值范围**: - handled : 已处理 - unhandled : 未处理  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 存在漏洞的容器ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 是否是容器场景 **约束限制**: 不涉及 **取值范围**: - true：是容器场景 - false：不是容器场景  **默认取值**: 不涉及
	IsContainer *bool `json:"is_container,omitempty"`
}

func (o ListVulHostAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostAppsRequest struct{}"
	}

	return strings.Join([]string{"ListVulHostAppsRequest", string(data)}, " ")
}
