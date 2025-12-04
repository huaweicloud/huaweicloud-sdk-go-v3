package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespPartitions struct {

	// **参数解释**： 分区名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespPartitions struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespPartitions", string(data)}, " ")
}
