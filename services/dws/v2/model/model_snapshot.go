package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Snapshot **参数解释**： 快照对象。 **取值范围**： 不涉及。
type Snapshot struct {

	// **参数解释**： 快照名称。要求唯一性且必须以字母开头，不区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符，长度为4～64个字符。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 指定创建快照的集群ID。 **取值范围**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 快照描述，若不指定，描述为空。快照描述的字符长度不能超过256个字符，且不支持特殊字符!<>'=&\"。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o Snapshot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Snapshot struct{}"
	}

	return strings.Join([]string{"Snapshot", string(data)}, " ")
}
