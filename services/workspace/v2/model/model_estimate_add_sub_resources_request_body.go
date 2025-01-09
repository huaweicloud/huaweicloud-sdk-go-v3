package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateAddSubResourcesRequestBody 包周期新增附属资源询价请求体。
type EstimateAddSubResourcesRequestBody struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 包周期桌面ID列表。 不可同时存在普通桌面和池桌面ID。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`
}

func (o EstimateAddSubResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateAddSubResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"EstimateAddSubResourcesRequestBody", string(data)}, " ")
}
