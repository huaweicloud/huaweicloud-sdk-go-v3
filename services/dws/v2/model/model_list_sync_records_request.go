package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncRecordsRequest Request Object
type ListSyncRecordsRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 1000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSyncRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListSyncRecordsRequest", string(data)}, " ")
}
