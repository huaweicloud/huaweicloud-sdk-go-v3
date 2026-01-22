package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVolumeExpandConfigResponse Response Object
type ShowVolumeExpandConfigResponse struct {

	// **参数解释**： 是否开启磁盘自动扩容。 **约束限制**： 不涉及。 **取值范围**： - true：开启。 - false：关闭。 **默认取值**： 不涉及。
	AutoVolumeExpandEnable *bool `json:"auto_volume_expand_enable,omitempty"`

	// **参数解释**： 触发磁盘自动扩容的阈值。 **约束限制**： 不涉及。 **取值范围**： 20%-80%。 **默认取值**： 不涉及。
	ExpandThreshold *int32 `json:"expand_threshold,omitempty"`

	// **参数解释**： 磁盘自动扩容的上限值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxVolumeSize *int32 `json:"max_volume_size,omitempty"`

	// **参数解释**： 每次磁盘自动扩容的比例。 **约束限制**： 不涉及。 **取值范围**： 10%-100%。 **默认取值**： 不涉及。
	ExpandIncrement *int32 `json:"expand_increment,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowVolumeExpandConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeExpandConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeExpandConfigResponse", string(data)}, " ")
}
