package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodePageInfo struct {

	// **参数解释**： 当前页返回的所有节点数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CurrentCount int32 `json:"currentCount"`

	// **参数解释**： 当前页最后一条记录，为查询下一页的marker取值，最后一页时无nextMarker字段。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NextMarker *string `json:"nextMarker,omitempty"`
}

func (o NodePageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePageInfo struct{}"
	}

	return strings.Join([]string{"NodePageInfo", string(data)}, " ")
}
