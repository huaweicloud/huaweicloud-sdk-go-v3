package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseUserRequest Request Object
type DeleteDatabaseUserRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 数据库用户名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 是否忽略大小写。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： false
	Cascade *bool `json:"cascade,omitempty"`
}

func (o DeleteDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserRequest", string(data)}, " ")
}
