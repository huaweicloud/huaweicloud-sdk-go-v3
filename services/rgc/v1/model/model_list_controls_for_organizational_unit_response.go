package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsForOrganizationalUnitResponse Response Object
type ListControlsForOrganizationalUnitResponse struct {

	// 治理策略概要。
	ControlSummaries *[]TargetControl `json:"control_summaries,omitempty"`

	// 控制策略启用状态。
	State *string `json:"state,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListControlsForOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsForOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"ListControlsForOrganizationalUnitResponse", string(data)}, " ")
}
