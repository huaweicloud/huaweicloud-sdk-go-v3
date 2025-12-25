package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicDov5 struct {

	// **参数解释**: 仓库状态。 **取值范围**: active：正常。 delete：删除。 disabled：无效。 view：私有库浏览者。 trash：废弃。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 租户id。 **取值范围**: 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**: 区域。 **取值范围**: 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**: 创建时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**: 不涉及。
	CreatedTime *string `json:"created_time,omitempty"`

	// **参数解释**: 修改时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**: 不涉及。
	ModifiedTime *string `json:"modified_time,omitempty"`

	// **参数解释**: 创建人id。 **取值范围**: 不涉及。
	CreatedUserId *string `json:"created_user_id,omitempty"`

	// **参数解释**: 创建人。 **取值范围**: 不涉及。
	CreatedUserName *string `json:"created_user_name,omitempty"`

	// **参数解释**: 修改人id。 **取值范围**: 不涉及。
	ModifiedUserId *string `json:"modified_user_id,omitempty"`

	// **参数解释**: 修改人。 **取值范围**: 不涉及。
	ModifiedUserName *string `json:"modified_user_name,omitempty"`
}

func (o BasicDov5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicDov5 struct{}"
	}

	return strings.Join([]string{"BasicDov5", string(data)}, " ")
}
