package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetaLabels 资源池metadata的标签信息。
type PoolMetaLabels struct {

	// **参数解释**：资源池的显示名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsName *string `json:"os.modelarts/name,omitempty"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc) **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	OsModelartsWorkspaceId *string `json:"os.modelarts/workspace.id,omitempty"`

	// **参数解释**：自定义节点前缀，可选值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsNodePrefix *string `json:"os.modelarts/node.prefix,omitempty"`

	// **参数解释**：资源池计费使用的资源ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsResourceId *string `json:"os.modelarts/resource.id,omitempty"`

	// **参数解释**：资源池所属的租户ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsTenantDomainId *string `json:"os.modelarts/tenant.domain.id,omitempty"`

	// **参数解释**：资源池所属的租户项目ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsTenantProjectId *string `json:"os.modelarts/tenant.project.id,omitempty"`

	// **参数解释**：资源池所属的企业项目ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsEnterpriseProjectId *string `json:"os.modelarts/enterprise.project.id,omitempty"`

	// **参数解释**：资源池商业类型，public是公共池，private个人专属池。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPoolBiz *string `json:"os.modelarts.pool/biz,omitempty"`

	// **参数解释**：资源池创建来源，比如admin-console，标记来自admin创建，console标记来自ma console。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsCreateFrom *string `json:"os.modelarts/create-from,omitempty"`

	// **参数解释**：资源池是否计费。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsNobilling *string `json:"os.modelarts/nobilling,omitempty"`

	// **参数解释**：资源池关联的上一次订单作业记录。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsOrderName *string `json:"os.modelarts/order.name,omitempty"`

	// **参数解释**：资源池所属区域。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsRegion *string `json:"os.modelarts/region,omitempty"`
}

func (o PoolMetaLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetaLabels struct{}"
	}

	return strings.Join([]string{"PoolMetaLabels", string(data)}, " ")
}
