package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSnapshotChainsResponse Response Object
type CountSnapshotChainsResponse struct {

	// **参数解释** 资源的数量。 **取值范围** 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountSnapshotChainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSnapshotChainsResponse struct{}"
	}

	return strings.Join([]string{"CountSnapshotChainsResponse", string(data)}, " ")
}
