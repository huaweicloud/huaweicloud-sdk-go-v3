package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDatabaseUsersRequest Request Object
type ExportDatabaseUsersRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 10000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 类型。 **约束限制**： 不涉及。 **取值范围**： ROLE：导出角色 USER：导出用户 **默认取值**： null
	Type *string `json:"type,omitempty"`
}

func (o ExportDatabaseUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDatabaseUsersRequest struct{}"
	}

	return strings.Join([]string{"ExportDatabaseUsersRequest", string(data)}, " ")
}
