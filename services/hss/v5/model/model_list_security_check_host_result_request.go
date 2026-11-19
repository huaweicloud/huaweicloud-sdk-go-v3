package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityCheckHostResultRequest Request Object
type ListSecurityCheckHostResultRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 操作系统类型 **约束限制**: 不涉及 **取值范围**: - Linux：Linux操作系统 - Windows：Windows操作系统  **默认取值**: 不涉及
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 风险级别 **约束限制**: 不涉及 **取值范围**: - Security：安全 - Low：低危 - Medium：中危 - High：高危  **默认取值**: 不涉及
	Severity *string `json:"severity,omitempty"`
}

func (o ListSecurityCheckHostResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityCheckHostResultRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityCheckHostResultRequest", string(data)}, " ")
}
