package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseProjectInfoResult struct {

	// **参数解释**: 企业项目ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 企业项目名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 企业项目描述。 **取值范围**: 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**: 企业项目状态。 **取值范围**: - 1：启用。 - 2：停用。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 创建时间，格式为UTC格式。如：2018-05-18T06:49:06Z。 **取值范围**: 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**: 修改时间，格式为UTC格式。如：2018-05-28T02:21:36Z。 **取值范围**: 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o EnterpriseProjectInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectInfoResult struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectInfoResult", string(data)}, " ")
}
