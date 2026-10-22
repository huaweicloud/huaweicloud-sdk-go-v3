package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotChainsResponse Response Object
type ListSnapshotChainsResponse struct {

	// **参数解释** 快照链列表。 **取值范围** 不涉及
	SnapshotChains *[]ListSnapshotChainsResponseBodySnapshotChains `json:"snapshot_chains,omitempty"`
	HttpStatusCode int                                             `json:"-"`
}

func (o ListSnapshotChainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotChainsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotChainsResponse", string(data)}, " ")
}
