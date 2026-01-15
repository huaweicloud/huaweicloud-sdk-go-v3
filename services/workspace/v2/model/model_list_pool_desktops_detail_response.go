package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolDesktopsDetailResponse Response Object
type ListPoolDesktopsDetailResponse struct {

	// 池桌面详情。
	PoolDesktops *[]PoolDesktopsDetailInfo `json:"pool_desktops,omitempty"`

	// 桌面总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 按需桌面总数。
	OnDemandDesktopsNum *int32 `json:"on_demand_desktops_num,omitempty"`

	// 包周期桌面总数。
	PeriodDesktopsNum *int32 `json:"period_desktops_num,omitempty"`

	// 按需免费桌面总数。
	OnDemandFreeImageDesktopsNum *int32 `json:"on_demand_free_image_desktops_num,omitempty"`

	// 按需收费桌面总数。
	OnDemandChargeImageDesktopsNum *int32 `json:"on_demand_charge_image_desktops_num,omitempty"`

	// 包周期免费桌面总数。
	PeriodFreeImageDesktopsNum *int32 `json:"period_free_image_desktops_num,omitempty"`

	// 包周期收费桌面总数。
	PeriodChargeImageDesktopsNum *int32 `json:"period_charge_image_desktops_num,omitempty"`

	InconsistentType *InconsistentTypeEnum `json:"inconsistent_type,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ListPoolDesktopsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolDesktopsDetailResponse struct{}"
	}

	return strings.Join([]string{"ListPoolDesktopsDetailResponse", string(data)}, " ")
}
