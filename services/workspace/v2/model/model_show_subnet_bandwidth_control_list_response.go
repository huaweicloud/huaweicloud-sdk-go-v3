package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubnetBandwidthControlListResponse Response Object
type ShowSubnetBandwidthControlListResponse struct {

	// 控制模式 - BLACK：黑名单控制。 - WHITE：白名单控制。
	ControlMode *string `json:"control_mode,omitempty"`

	// 控制的桌面列表。
	ControlList *[]ControlItem `json:"control_list,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSubnetBandwidthControlListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubnetBandwidthControlListResponse struct{}"
	}

	return strings.Join([]string{"ShowSubnetBandwidthControlListResponse", string(data)}, " ")
}
