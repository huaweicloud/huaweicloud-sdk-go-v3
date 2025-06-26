package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUsersRequest Request Object
type ListDatabaseUsersRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 1000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 查询角色还是用户。 **约束限制**： 不涉及。 **取值范围**： ROLE：表示查询所有角色。  USER：表示查询所有用户。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 用户类型，COMMON、IAM或者OneAccess。 **约束限制**： 不涉及。 **取值范围**： COMMON：表示普通数据库用户。  IAM：表示IAM同步的数据库用户。 OneAccess: 表示OneAccess用户。 **默认取值**： 不涉及。
	UserType *string `json:"user_type,omitempty"`
}

func (o ListDatabaseUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseUsersRequest", string(data)}, " ")
}
