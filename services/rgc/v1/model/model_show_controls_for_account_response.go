package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlsForAccountResponse Response Object
type ShowControlsForAccountResponse struct {
	Control *EnabledControl `json:"control,omitempty"`

	// 区域信息。
	Regions *[]RegionConfigurationList `json:"regions,omitempty"`

	// 状态。
	State *string `json:"state,omitempty"`

	// 状态说明。
	Message *string `json:"message,omitempty"`

	// 控制策略当前版本。
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowControlsForAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlsForAccountResponse struct{}"
	}

	return strings.Join([]string{"ShowControlsForAccountResponse", string(data)}, " ")
}
