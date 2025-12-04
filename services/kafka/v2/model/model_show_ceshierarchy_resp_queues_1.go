package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespQueues1 struct {

	// **参数解释**： Topic名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  分区信息。
	Partitions *[]ShowCeshierarchyRespPartitions `json:"partitions,omitempty"`
}

func (o ShowCeshierarchyRespQueues1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespQueues1 struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespQueues1", string(data)}, " ")
}
