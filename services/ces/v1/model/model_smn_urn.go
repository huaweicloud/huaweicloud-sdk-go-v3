package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnUrn **参数解释**： 通知对象ID。 **约束限制**： 不涉及。 **取值范围**： 只能包含字母、数字、“-”、“_” 、“:”。     **默认取值**： 不涉及。
type SmnUrn struct {
}

func (o SmnUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnUrn struct{}"
	}

	return strings.Join([]string{"SmnUrn", string(data)}, " ")
}
