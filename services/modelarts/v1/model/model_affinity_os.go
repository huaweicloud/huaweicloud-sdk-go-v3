package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AffinityOs **参数解释**：Modelarts内置操作系统。
type AffinityOs struct {

	// **参数解释**：操作系统名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：操作系统镜像id, 裸金属规格该字段不为空。 **取值范围**：不涉及。
	ImageId *string `json:"imageId,omitempty"`

	// **参数解释**：是否优选。 **取值范围**：不涉及。
	Preferred *bool `json:"preferred,omitempty"`

	// **参数解释**：操作系统是否即将停止服务, end of service。 **取值范围**：不涉及。
	Eos *bool `json:"eos,omitempty"`

	// **参数解释**：操作系统是否下线。 **取值范围**：不涉及
	Offline *bool `json:"offline,omitempty"`
}

func (o AffinityOs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AffinityOs struct{}"
	}

	return strings.Join([]string{"AffinityOs", string(data)}, " ")
}
