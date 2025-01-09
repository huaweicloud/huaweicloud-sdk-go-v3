package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInquiryCommonInfo 批量询价/下单公共参数
type BatchInquiryCommonInfo struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 包周期桌面ID列表。 不可同时存在普通桌面和池桌面ID。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`
}

func (o BatchInquiryCommonInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInquiryCommonInfo struct{}"
	}

	return strings.Join([]string{"BatchInquiryCommonInfo", string(data)}, " ")
}
