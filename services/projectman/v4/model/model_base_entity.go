package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseEntity 通用实体信息
type BaseEntity struct {

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释**： 修改人。 **取值范围**： 不涉及。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// **参数解释**： 修改时间。 **取值范围**： 不涉及。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// **参数解释**： 创建人。 **取值范围**： 不涉及。
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreatedDate *string `json:"created_date,omitempty"`
}

func (o BaseEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseEntity struct{}"
	}

	return strings.Join([]string{"BaseEntity", string(data)}, " ")
}
