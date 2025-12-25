package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteClusterUpgradeActionRequestBody struct {

	// **参数解释**： 当前集群要做的操作。 **约束限制**： 不涉及。 **取值范围**： - update: 下发升级 - rollback: 下发回滚 - commit: 下发提交集群 - retry: 下发重试 **默认取值**： 不涉及。
	Action string `json:"action"`

	// **参数解释**： 升级项ID。 **约束限制**： 填写的升级项ID应属于当前集群。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ItemId string `json:"item_id"`
}

func (o ExecuteClusterUpgradeActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteClusterUpgradeActionRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteClusterUpgradeActionRequestBody", string(data)}, " ")
}
