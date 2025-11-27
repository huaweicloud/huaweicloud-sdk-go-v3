package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncResourceAgentReqResourceInfos struct {

	// **参数解释：** 资源所属的region id。 **约束限制：** 不涉及。 **取值范围：** 资源对应区域id。 **默认取值：** 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 资源id。 **约束限制：** 不涉及。 **取值范围：** 用户选择的资源对应的资源id。 **默认取值：** 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释：** 资源在coc-cmdb存储的id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`
}

func (o SyncResourceAgentReqResourceInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourceAgentReqResourceInfos struct{}"
	}

	return strings.Join([]string{"SyncResourceAgentReqResourceInfos", string(data)}, " ")
}
