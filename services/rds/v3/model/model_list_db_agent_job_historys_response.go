package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbAgentJobHistorysResponse Response Object
type ListDbAgentJobHistorysResponse struct {

	// 执行历史列表。
	Historys *[]ListDbAgentJobHistorysResult `json:"historys,omitempty"`

	// 执行历史总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbAgentJobHistorysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbAgentJobHistorysResponse struct{}"
	}

	return strings.Join([]string{"ListDbAgentJobHistorysResponse", string(data)}, " ")
}
