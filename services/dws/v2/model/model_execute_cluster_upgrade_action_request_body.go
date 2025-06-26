package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteClusterUpgradeActionRequestBody struct {

	// **参数解释**： 当前集群要做的操作。 **取值范围**： 不涉及。
	Action string `json:"action"`

	// **参数解释**： 升级项ID。 **取值范围**： 不涉及。
	ItemId string `json:"item_id"`
}

func (o ExecuteClusterUpgradeActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteClusterUpgradeActionRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteClusterUpgradeActionRequestBody", string(data)}, " ")
}
