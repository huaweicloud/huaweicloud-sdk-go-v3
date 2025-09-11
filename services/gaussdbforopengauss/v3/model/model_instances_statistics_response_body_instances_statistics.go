package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstancesStatisticsResponseBodyInstancesStatistics struct {

	// **参数解释**: 实例状态。 **取值范围**: - normal：实例状态正常。 - abnormal：实例状态异常。 - creating：实例创建中。 - createfail：实例创建失败。 - shutdown：实例已关机。 - deleted：实例已删除。
	Status string `json:"status"`

	// **参数解释**: 实例数量。 **取值范围**: 不涉及。
	Count int32 `json:"count"`
}

func (o InstancesStatisticsResponseBodyInstancesStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesStatisticsResponseBodyInstancesStatistics struct{}"
	}

	return strings.Join([]string{"InstancesStatisticsResponseBodyInstancesStatistics", string(data)}, " ")
}
