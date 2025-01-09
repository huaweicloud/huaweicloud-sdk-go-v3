package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resource 包周期资源对象
type Resource struct {

	// 订购周期类型：2：月；3：年；4：包小时（仅限带宽加油包购买场景使用）5：绝对时间；（追加附属资源场景使用，比如主机上追加云硬盘）6：一次性（chargingMode=1 一次性计费场景使用），必填
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购周期数
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否续订
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	AddVolumes *AddVolumes `json:"add_volumes,omitempty"`

	CreateDesktops *CreateDesktopReq `json:"create_desktops,omitempty"`

	DehHosts *Hosts `json:"deh_hosts,omitempty"`

	RebuildDesktops *RebuildDesktopsReq `json:"rebuild_desktops,omitempty"`

	AttachDesktops *AttachInstancesReq `json:"attach_desktops,omitempty"`

	CreateExclusiveHosts *CreateExclusiveHostsReq `json:"create_exclusive_hosts,omitempty"`

	ResizeExclusiveLites *ResizeExclusiveLitesReq `json:"resize_exclusive_lites,omitempty"`

	CreateDesktopPool *CreateDesktopPoolReq `json:"create_desktop_pool,omitempty"`

	ExpandDesktopPool *ExpandDesktopPoolOrderReq `json:"expand_desktop_pool,omitempty"`

	ApplyDesktopsInternet *ApplyDesktopsInternet `json:"apply_desktops_internet,omitempty"`

	ApplySubnetBandwidth *ApplySubnetBandwidthReq `json:"apply_subnet_bandwidth,omitempty"`

	SubscribeUserSharer *SubscribeUserSharerReq `json:"subscribe_user_sharer,omitempty"`

	// 支付后跳转的地址
	CloudServiceConsoleUrl *string `json:"cloud_service_console_url,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
