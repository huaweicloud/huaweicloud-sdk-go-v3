package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListViewResponseBodyData struct {

	// **参数解释：** 资源视图id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 资源视图名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 视图创建者的domainId，视图资源聚合的归属者。 **取值范围：** 不涉及。
	ManagerDomainId *string `json:"manager_domain_id,omitempty"`

	// **参数解释：** 视图归属的组织id。 **取值范围：** 不涉及。
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释：** 视图所聚合的组织单元id列表。 **取值范围：** 不涉及。
	OrganizationUnitIds *[]string `json:"organization_unit_ids,omitempty"`

	// **参数解释：** 视图包含的租户账号id列表。 **取值范围：** 不涉及。
	DomainIds *[]string `json:"domain_ids,omitempty"`

	// **参数解释：** 视图包含的资源类型列表。 **取值范围：** 不涉及。
	ResourceTypes *[]string `json:"resource_types,omitempty"`

	// **参数解释：** 视图类型。 **取值范围：** 不涉及。
	ViewType *string `json:"view_type,omitempty"`
}

func (o ListViewResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListViewResponseBodyData struct{}"
	}

	return strings.Join([]string{"ListViewResponseBodyData", string(data)}, " ")
}
