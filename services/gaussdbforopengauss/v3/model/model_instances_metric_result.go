package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstancesMetricResult struct {

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 实例状态。 **取值范围**： - creating，表示实例正在创建。 - normal，表示实例正常。 - abnormal，表示实例异常。 - createfail，表示实例创建失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 实例类型。 **取值范围**： 不涉及。
	Mode *string `json:"mode,omitempty"`

	// **参数解释**： 引擎名称。 **取值范围**： 不涉及。
	EngineName *string `json:"engine_name,omitempty"`

	// **参数解释**： 引擎版本。 **取值范围**： 不涉及。
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**： 部署形态。 **取值范围**： 不涉及。
	Solution *string `json:"solution,omitempty"`

	// **参数解释**： 实例数据磁盘已使用大小。 **取值范围**： 不涉及。
	DiskUsedSize *string `json:"disk_used_size,omitempty"`

	// **参数解释**： 实例数据磁盘总大小。 **取值范围**： 不涉及。
	DiskTotalSize *string `json:"disk_total_size,omitempty"`

	// **参数解释**： 实例数据磁盘已使用百分比。 **取值范围**： 不涉及。
	DiskUsage *string `json:"disk_usage,omitempty"`

	// **参数解释**： 80% SQL的响应时间。 **取值范围**： 不涉及。
	P80 *string `json:"p80,omitempty"`

	// **参数解释**： 95% SQL的响应时间。 **取值范围**： 不涉及。
	P95 *string `json:"p95,omitempty"`

	// **参数解释**： 死锁次数。 **取值范围**： 不涉及。
	Deadlocks *string `json:"deadlocks,omitempty"`

	// **参数解释**： buffer 命中率。 **取值范围**： 不涉及。
	BufferHitRatio *string `json:"buffer_hit_ratio,omitempty"`

	// **参数解释**： 实例节点信息列表。
	Nodes *[]InstancesNodesResult `json:"nodes,omitempty"`
}

func (o InstancesMetricResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesMetricResult struct{}"
	}

	return strings.Join([]string{"InstancesMetricResult", string(data)}, " ")
}
