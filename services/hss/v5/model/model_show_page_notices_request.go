package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPageNoticesRequest Request Object
type ShowPageNoticesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 访问页面位置，包含下面这些页面 - hostMgmt : 主机管理-云服务器 - hostProtectQuota : 主机管理-防护配额 - containerNodeList : 容器管理-容器节点管理 - containerProtectQuota : 容器管理-容器防护配额 - containerMirror : 容器管理-容器镜像 - container : 容器管理-容器 - clusterAgent : 容器管理-集群Agent管理 - vulView : 漏洞管理-漏洞视图 - vulHostView : 漏洞管理-主机视图 - ransomwareProtection : 勒索病毒防护 - policyMgmt : 策略管理 - antiVirus : 病毒查杀 - hostAlarm : 安全告警事件-主机安全告警 - containerAlarm : 安全告警事件-容器安全告警
	PageLocation *string `json:"page_location,omitempty"`
}

func (o ShowPageNoticesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPageNoticesRequest struct{}"
	}

	return strings.Join([]string{"ShowPageNoticesRequest", string(data)}, " ")
}
