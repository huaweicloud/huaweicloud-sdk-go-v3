package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountVo struct {

	// **参数解释**： 账号ID **取值范围**： 不涉及
	AccountId *string `json:"account_id,omitempty"`

	// **参数解释**： 账号名称 **取值范围**： 不涉及
	AccountName *string `json:"account_name,omitempty"`

	// **参数解释**： 防护的EIP数量 **取值范围**： 不涉及
	EipCountProtected *int32 `json:"eip_count_protected,omitempty"`

	// **参数解释**： EIP总数 **取值范围**： 不涉及
	EipCountTotal *int32 `json:"eip_count_total,omitempty"`

	// **参数解释**： 未开启防护的EIP数量 **取值范围**： 不涉及
	EipCountUnprotected *int32 `json:"eip_count_unprotected,omitempty"`

	// **参数解释**： 组织ID **取值范围**： 不涉及
	OrganizationId *string `json:"organization_id,omitempty"`

	// **参数解释**： 项目ID **取值范围**： 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 项目名称 **取值范围**： 不涉及
	ProjectName *string `json:"project_name,omitempty"`
}

func (o AccountVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountVo struct{}"
	}

	return strings.Join([]string{"AccountVo", string(data)}, " ")
}
