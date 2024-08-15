package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScriptOrderStatisticsModel 任务统计信息  讨论： statistics属性里用状态声明key合理么？ 还要统计什么内容，怎么扩展？
type JobScriptOrderStatisticsModel struct {

	// 实例总量
	TotalInstance *int32 `json:"total_instance,omitempty"`

	// 每个状态一个count，里面记录该状态的总数量，以及包含该状态的批次列表
	ExecuteStatistics *[]ExectuionStatistic `json:"execute_statistics,omitempty"`
}

func (o JobScriptOrderStatisticsModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptOrderStatisticsModel struct{}"
	}

	return strings.Join([]string{"JobScriptOrderStatisticsModel", string(data)}, " ")
}
