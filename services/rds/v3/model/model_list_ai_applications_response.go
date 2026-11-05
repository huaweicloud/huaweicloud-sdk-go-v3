package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiApplicationsResponse Response Object
type ListAiApplicationsResponse struct {
	TotalCount *int64 `json:"total_count,omitempty"`

	Data           *[]AiApplicationListItem `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAiApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListAiApplicationsResponse", string(data)}, " ")
}
