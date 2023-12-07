package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlsForOrganizationUnitResponse Response Object
type ShowControlsForOrganizationUnitResponse struct {
	Control *EnabledControl `json:"control,omitempty"`

	// 开启的区域。
	Regions *[]string `json:"regions,omitempty"`

	// 是否允许启用控制策略。
	State *string `json:"state,omitempty"`

	// 状态说明。
	StateMessage *string `json:"state_message,omitempty"`

	// 控制策略当前版本。
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowControlsForOrganizationUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlsForOrganizationUnitResponse struct{}"
	}

	return strings.Join([]string{"ShowControlsForOrganizationUnitResponse", string(data)}, " ")
}
