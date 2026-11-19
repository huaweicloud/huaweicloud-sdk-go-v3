package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppliedHistoriesWithEntitiesResponse Response Object
type ListAppliedHistoriesWithEntitiesResponse struct {

	// 应用记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 应用记录列表。
	Histories      *[]AppliedHistoriesResult `json:"histories,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAppliedHistoriesWithEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppliedHistoriesWithEntitiesResponse struct{}"
	}

	return strings.Join([]string{"ListAppliedHistoriesWithEntitiesResponse", string(data)}, " ")
}
