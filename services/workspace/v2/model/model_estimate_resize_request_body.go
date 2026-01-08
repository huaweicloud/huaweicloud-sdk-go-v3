package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EstimateResizeRequestBody struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 包周期桌面ID列表。 不可同时存在普通桌面和池桌面ID。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *string `json:"handle_type,omitempty"`
}

func (o EstimateResizeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateResizeRequestBody struct{}"
	}

	return strings.Join([]string{"EstimateResizeRequestBody", string(data)}, " ")
}
