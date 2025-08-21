package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OrganizationAccountInfo struct {

	// **参数解释**： 是否已添加 **取值范围**： 不涉及
	Delegated *bool `json:"delegated,omitempty"`

	// **参数解释**： 账号的唯一标识符 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 账号名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 组织节点类型 **取值范围**： account
	OrgType *string `json:"org_type,omitempty"`

	// **参数解释**： 父节点（根或组织单元）的唯一标识符（ID） **取值范围**： 不涉及
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**： 账号的统一资源名称 **取值范围**： 不涉及
	Urn *string `json:"urn,omitempty"`
}

func (o OrganizationAccountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationAccountInfo struct{}"
	}

	return strings.Join([]string{"OrganizationAccountInfo", string(data)}, " ")
}
