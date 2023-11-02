package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteClusterUpgradeActionRequestBody struct {

	// 当前集群要做的操作
	Action string `json:"action"`

	// 升级项ID
	ItemId string `json:"item_id"`
}

func (o ExecuteClusterUpgradeActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteClusterUpgradeActionRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteClusterUpgradeActionRequestBody", string(data)}, " ")
}
