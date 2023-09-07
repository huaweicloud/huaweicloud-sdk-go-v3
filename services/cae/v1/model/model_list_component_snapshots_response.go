package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentSnapshotsResponse Response Object
type ListComponentSnapshotsResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentSnapshotKindObj `json:"kind,omitempty"`

	// 快照列表。
	Items          *[]ComponentSnapshotItem `json:"items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListComponentSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListComponentSnapshotsResponse", string(data)}, " ")
}
