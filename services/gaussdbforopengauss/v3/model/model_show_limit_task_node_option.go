package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowLimitTaskNodeOption struct {

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 该节点执行的sql语句ID。 **取值范围**: 不涉及。
	SqlId string `json:"sql_id"`
}

func (o ShowLimitTaskNodeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLimitTaskNodeOption struct{}"
	}

	return strings.Join([]string{"ShowLimitTaskNodeOption", string(data)}, " ")
}
