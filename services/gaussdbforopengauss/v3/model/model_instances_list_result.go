package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstancesListResult struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o InstancesListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesListResult struct{}"
	}

	return strings.Join([]string{"InstancesListResult", string(data)}, " ")
}
