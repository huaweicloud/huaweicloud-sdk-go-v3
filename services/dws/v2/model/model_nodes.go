package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nodes **参数解释**： 集群实例对象。 **取值范围**： 非空对象列表。
type Nodes struct {

	// **参数解释**： 集群实例ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 集群实例状态。 **取值范围**： - 100：创建中 - 199：空闲 - 200：可用 - 300：不可用 - 303：创建失败 - 304：删除中 - 305：删除失败 - 400：已删除
	Status string `json:"status"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o Nodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nodes struct{}"
	}

	return strings.Join([]string{"Nodes", string(data)}, " ")
}
