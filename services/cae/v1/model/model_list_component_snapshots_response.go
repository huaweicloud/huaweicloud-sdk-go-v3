package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentSnapshotsResponse Response Object
type ListComponentSnapshotsResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“ComponentSnapshot”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
