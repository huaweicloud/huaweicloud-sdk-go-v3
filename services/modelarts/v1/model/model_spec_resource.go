package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecResource 训练作业资源规格信息。flavor_id和pool_id+[flavor_id]方式二选一。 - 选择公共资源池时，仅上送flavor_id，选择训练作业需要的卡数、内存等资源规格，当公共资源池空闲资源满足选择的规格需求时，作业可被调度； - 选择专属资源池时，需上送pool_id与flavor_id，选择专属资源池下可选的实际规格，即满足训练作业条件的最小卡数，以便节省专属资源，提高利用率。
type SpecResource struct {

	// **参数解释**：训练作业资源规格id。 **约束限制**：不涉及。 **取值范围**：CPU规格专属资源池不支持指定flavor_id。GPU/Ascend规格专属资源池可选取值如下： - modelarts.pool.visual.xlarge（1卡） - modelarts.pool.visual.2xlarge（2卡） - modelarts.pool.visual.4xlarge（4卡） - modelarts.pool.visual.8xlarge（8卡） - modelarts.pool.visual.16xlarge（16卡，当前仅限Snt9b23超节点资源池）  **默认取值**：不涉及。
	FlavorId *string `json:"flavor_id,omitempty"`

	// **参数解释**：资源池创建训练作业使用节点数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：默认单节点。
	NodeCount int32 `json:"node_count"`

	// **参数解释**：专属资源池id。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释**：资源池联邦id。 **约束限制**：当kind为federated_pool_job时，该字段必填。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolGroupId *string `json:"pool_group_id,omitempty"`

	MainContainerCustomizedFlavor *MainContainerCustomizedFlavor `json:"main_container_customized_flavor,omitempty"`
}

func (o SpecResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecResource struct{}"
	}

	return strings.Join([]string{"SpecResource", string(data)}, " ")
}
