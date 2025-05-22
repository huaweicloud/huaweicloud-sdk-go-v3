package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterVolume **参数解释**： 逻辑集群磁盘信息。 **取值范围**： 不涉及。
type LogicalClusterVolume struct {

	// **参数解释**： 逻辑集群名称。 **取值范围**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 磁盘使用量。 **取值范围**： 不涉及。
	Usage *string `json:"usage,omitempty"`

	// **参数解释**： 磁盘总量。 **取值范围**： 不涉及。
	Total *string `json:"total,omitempty"`

	// **参数解释**： 磁盘使用比例。 **取值范围**： 不涉及。
	Percent *string `json:"percent,omitempty"`
}

func (o LogicalClusterVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterVolume struct{}"
	}

	return strings.Join([]string{"LogicalClusterVolume", string(data)}, " ")
}
