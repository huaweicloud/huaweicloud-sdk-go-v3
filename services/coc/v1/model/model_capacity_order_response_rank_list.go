package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CapacityOrderResponseRankList **参数解释：** TOP5排名对象列表。 **取值范围：** 容量种类排名前五的应用或组件或分组的容量信息组成的列表。
type CapacityOrderResponseRankList struct {

	// **参数解释：** 应用或组件或者分组id。 **取值范围：** 容量种类排名在前五个的应用或者组件或分组对应的id。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 应用或组件或者分组的名称。 **取值范围：** 容量种类排名在前五个的应用或者组件或分组对应的名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 应用或组件或者分组的容量值。 **取值范围：** 容量种类排名在前五个的应用或者组件或分组对应的容量值。
	Value *string `json:"value,omitempty"`
}

func (o CapacityOrderResponseRankList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityOrderResponseRankList struct{}"
	}

	return strings.Join([]string{"CapacityOrderResponseRankList", string(data)}, " ")
}
