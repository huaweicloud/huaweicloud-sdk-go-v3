package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsForOrganizationUnitResponse Response Object
type ListControlsForOrganizationUnitResponse struct {

	// 治理策略概要。
	ControlSummaries *[]TargetControl `json:"control_summaries,omitempty"`

	// 控制策略启用状态。
	State *string `json:"state,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListControlsForOrganizationUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsForOrganizationUnitResponse struct{}"
	}

	return strings.Join([]string{"ListControlsForOrganizationUnitResponse", string(data)}, " ")
}
