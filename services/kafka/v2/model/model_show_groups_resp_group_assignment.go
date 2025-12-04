package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowGroupsRespGroupAssignment struct {

	// **参数解释**： Topic名称。 **取值范围**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 分区列表。
	Partitions *[]int32 `json:"partitions,omitempty"`
}

func (o ShowGroupsRespGroupAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupAssignment struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupAssignment", string(data)}, " ")
}
