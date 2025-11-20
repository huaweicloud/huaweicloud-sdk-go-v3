package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobSpec
type JobSpec struct {

	// **参数解释**： 任务的类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 任务所在的集群的ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClusterUID *string `json:"clusterUID,omitempty"`

	// **参数解释**： 任务操作的资源ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ResourceID *string `json:"resourceID,omitempty"`

	// **参数解释**： 任务操作的资源名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ResourceName *string `json:"resourceName,omitempty"`

	// **参数解释**： 扩展参数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ExtendParam map[string]string `json:"extendParam,omitempty"`

	// **参数解释**： 子任务的列表。 - 包含了所有子任务的详细信息 - 在创建集群、节点等场景下，通常会由多个子任务共同组成创建任务，在子任务都完成后，任务才会完成  **约束限制**： 不涉及
	SubJobs *[]Job `json:"subJobs,omitempty"`
}

func (o JobSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSpec struct{}"
	}

	return strings.Join([]string{"JobSpec", string(data)}, " ")
}
