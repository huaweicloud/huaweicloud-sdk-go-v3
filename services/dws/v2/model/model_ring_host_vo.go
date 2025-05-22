package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RingHostVo **参数解释**：  集群环信息。 **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
type RingHostVo struct {

	// **参数解释**：  主机名。 **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	HostName string `json:"host_name"`
}

func (o RingHostVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RingHostVo struct{}"
	}

	return strings.Join([]string{"RingHostVo", string(data)}, " ")
}
