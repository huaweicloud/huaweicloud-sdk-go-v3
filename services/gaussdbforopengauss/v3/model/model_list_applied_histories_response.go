package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppliedHistoriesResponse Response Object
type ListAppliedHistoriesResponse struct {

	// 应用记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 应用记录列表。
	Histories      *[]AppliedHistoriesResult `json:"histories,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAppliedHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppliedHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAppliedHistoriesResponse", string(data)}, " ")
}
