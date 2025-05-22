package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateAzRequestBody struct {

	// **参数解释**： 新可用区。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NewAvailableZones *[]string `json:"new_available_zones,omitempty"`

	// **参数解释**： 是否立即执行变更。 **约束限制**： 不涉及。 **取值范围**： true：立即执行变更。 false：暂不执行变更。 **默认取值**： true
	ExecuteImmediately *bool `json:"execute_immediately,omitempty"`
}

func (o MigrateAzRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzRequestBody struct{}"
	}

	return strings.Join([]string{"MigrateAzRequestBody", string(data)}, " ")
}
