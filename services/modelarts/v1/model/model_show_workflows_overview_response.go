package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowsOverviewResponse Response Object
type ShowWorkflowsOverviewResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 状态。
	Stat           map[string]int32 `json:"stat,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowWorkflowsOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowsOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowsOverviewResponse", string(data)}, " ")
}
