package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUserAuthoritiesRequest Request Object
type ListDatabaseUserAuthoritiesRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 用户名、角色名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseUserAuthoritiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUserAuthoritiesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseUserAuthoritiesRequest", string(data)}, " ")
}
