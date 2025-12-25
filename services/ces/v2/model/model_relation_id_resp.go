package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelationIdResp **参数解释**： 关联ID       **取值范围**： 取值为告警规则ID、告警策略ID。只能包含字母、数字、“-”，长度为[1,64]个字符。
type RelationIdResp struct {
}

func (o RelationIdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationIdResp struct{}"
	}

	return strings.Join([]string{"RelationIdResp", string(data)}, " ")
}
