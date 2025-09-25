package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteSqlLimitTaskRequestBody struct {

	// **参数解释**: 限流任务所在的节点ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`
}

func (o DeleteSqlLimitTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitTaskRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitTaskRequestBody", string(data)}, " ")
}
