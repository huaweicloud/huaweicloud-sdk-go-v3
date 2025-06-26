package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotFlavorInfoResponse Response Object
type ListSnapshotFlavorInfoResponse struct {

	// **参数解释**： 快照规格信息响应。 **取值范围**： 不涉及。
	Flavors *[]ProductUnitResp `json:"flavors,omitempty"`

	// **参数解释**： 快照规格信息总条数。 **取值范围**： 大于等于0的正整数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSnapshotFlavorInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotFlavorInfoResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotFlavorInfoResponse", string(data)}, " ")
}
