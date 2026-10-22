package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSnapshotsGroupResponse Response Object
type CountSnapshotsGroupResponse struct {

	// **参数解释** 资源的数量。 **取值范围** 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountSnapshotsGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSnapshotsGroupResponse struct{}"
	}

	return strings.Join([]string{"CountSnapshotsGroupResponse", string(data)}, " ")
}
