package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResizeOrderRequestBody 包周期变更规格询价请求体。
type CreateResizeOrderRequestBody struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 桌面id
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *string `json:"handle_type,omitempty"`

	// 目标规格产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 是否支持开机状态下执行变更规格操作。固定传值STOP_DESKTOP，如果桌面处于开机状态，会先关机再变更规格。
	Mode *string `json:"mode,omitempty"`
}

func (o CreateResizeOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResizeOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResizeOrderRequestBody", string(data)}, " ")
}
