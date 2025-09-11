package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelationId **参数解释**： 关联ID       **约束限制**： 不涉及。 **取值范围**： 取值为告警规则ID、告警策略ID。只能包含字母、数字、“-”，长度为[1,64]个字符。      **默认取值**： 不涉及。
type RelationId struct {
}

func (o RelationId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationId struct{}"
	}

	return strings.Join([]string{"RelationId", string(data)}, " ")
}
