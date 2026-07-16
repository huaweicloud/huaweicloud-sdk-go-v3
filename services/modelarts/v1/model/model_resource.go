package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resource 训练作业资源规格信息。flavor_id和pool_id+[flavor_id]方式二选一。
type Resource struct {

	// **参数解释**：训练作业资源规格模式。 **取值范围**： - regular：标准模式
	Policy *string `json:"policy,omitempty"`

	// **参数解释**：训练作业资源规格id。 **取值范围**：CPU规格专属资源池不支持指定flavor_id。GPU/Ascend规格专属资源池可选取值如下： - modelarts.pool.visual.xlarge（1卡） - modelarts.pool.visual.2xlarge（2卡） - modelarts.pool.visual.4xlarge（4卡） - modelarts.pool.visual.8xlarge（8卡）
	FlavorId string `json:"flavor_id"`

	// **参数解释**：使用flavor_id时，由ModelArts返回的只读规格名称。 **取值范围**：不涉及。
	FlavorName *string `json:"flavor_name,omitempty"`

	// **参数解释**：训练作业选择的资源副本数。 **取值范围**：不涉及。
	NodeCount int32 `json:"node_count"`

	// **参数解释**：训练作业选择的资源池ID。 **取值范围**：不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释**：训练作业选择的资源池联邦ID。 **取值范围**：不涉及。
	PoolGroupId *string `json:"pool_group_id,omitempty"`

	FlavorDetail *FlavorDetail `json:"flavor_detail,omitempty"`

	MainContainerAllocatedResources *MainContainerAllocatedResources `json:"main_container_allocated_resources,omitempty"`

	MainContainerCustomizedFlavor *MainContainerCustomizedFlavor `json:"main_container_customized_flavor,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
