package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceId **参数解释**： 实例ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type InstanceId struct {
}

func (o InstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceId struct{}"
	}

	return strings.Join([]string{"InstanceId", string(data)}, " ")
}
