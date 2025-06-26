package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseUserRequest Request Object
type ShowDatabaseUserRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 用户名/角色名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`
}

func (o ShowDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseUserRequest", string(data)}, " ")
}
