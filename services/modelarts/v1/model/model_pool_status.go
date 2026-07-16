package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolStatus 资源池状态信息。
type PoolStatus struct {

	// **参数解释**：资源池的状态。 **取值范围**：可选值如下： - Creating：资源池在创建中。 - Running：资源池在运行中。 - Abnormal：资源池异常。 - Deleting：资源池在删除中。 - Error：资源池错误。
	Phase string `json:"phase"`

	// **参数解释**：资源池当前状态的提示信息。 **取值范围**：不涉及。
	Message *string `json:"message,omitempty"`

	Resources *PoolStatusResources `json:"resources,omitempty"`

	// **参数解释**：资源池当前支持的业务类型的状态信息。
	Scope *[]PoolStatusScope `json:"scope,omitempty"`

	Driver *PoolStatusDriver `json:"driver,omitempty"`

	// **参数解释**：资源池所属父资源池的ID。物理池为空。 **取值范围**：不涉及。
	Parent *string `json:"parent,omitempty"`

	// **参数解释**：资源池根资源池的ID。 **取值范围**：不涉及。
	Root *string `json:"root,omitempty"`

	Clusters *PoolStatusClusters `json:"clusters,omitempty"`
}

func (o PoolStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatus struct{}"
	}

	return strings.Join([]string{"PoolStatus", string(data)}, " ")
}
