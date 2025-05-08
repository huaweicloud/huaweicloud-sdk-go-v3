package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterPlanBo 更新逻辑集群增删计划
type UpdateLogicalClusterPlanBo struct {

	// 更新逻辑集群增删计划细节信息列表
	Actions []UpdateLogicalClusterPlanActions `json:"actions"`
}

func (o UpdateLogicalClusterPlanBo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterPlanBo struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterPlanBo", string(data)}, " ")
}
