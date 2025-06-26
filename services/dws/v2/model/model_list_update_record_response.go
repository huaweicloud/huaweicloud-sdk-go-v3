package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdateRecordResponse Response Object
type ListUpdateRecordResponse struct {

	// **参数解释**： 集群升级记录总数。 **取值范围**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 集群升级记录列表。 **取值范围**： 不涉及。
	Records        *[]ClusterUpdateRecordResp `json:"records,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListUpdateRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdateRecordResponse struct{}"
	}

	return strings.Join([]string{"ListUpdateRecordResponse", string(data)}, " ")
}
