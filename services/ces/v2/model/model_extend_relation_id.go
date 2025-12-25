package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtendRelationId **参数解释** 关联企业项目ID。 **约束限制** 不涉及。 **取值范围** 只能包含字母、数字、和-，长度为36个字符；或者为0，代表默认企业项目 **默认取值** 不涉及。
type ExtendRelationId struct {
}

func (o ExtendRelationId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendRelationId struct{}"
	}

	return strings.Join([]string{"ExtendRelationId", string(data)}, " ")
}
