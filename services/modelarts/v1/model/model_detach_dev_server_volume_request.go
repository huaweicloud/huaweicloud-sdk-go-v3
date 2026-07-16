package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachDevServerVolumeRequest Request Object
type DetachDevServerVolumeRequest struct {

	// **参数解释**：Lite Server实例ID。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`

	// **参数解释**：要卸载的磁盘ID。 **约束限制**：[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}。 **取值范围**：不涉及。 **默认取值**：不涉及。
	VolumeId string `json:"volume_id"`
}

func (o DetachDevServerVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachDevServerVolumeRequest struct{}"
	}

	return strings.Join([]string{"DetachDevServerVolumeRequest", string(data)}, " ")
}
