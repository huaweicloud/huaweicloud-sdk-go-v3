package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateAzRequest Request Object
type MigrateAzRequest struct {

	// **参数解释**： 实例ID。可通过DCS控制台进入实例详情界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	Body *MigrateAzRequestBody `json:"body,omitempty"`
}

func (o MigrateAzRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzRequest struct{}"
	}

	return strings.Join([]string{"MigrateAzRequest", string(data)}, " ")
}
