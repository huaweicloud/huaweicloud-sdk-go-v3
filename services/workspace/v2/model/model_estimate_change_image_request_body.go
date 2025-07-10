package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateChangeImageRequestBody 包周期桌面池由不收费镜像切换至收费镜像的询价请求体。
type EstimateChangeImageRequestBody struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 包周期桌面ID列表。 不可同时存在普通桌面和池桌面ID。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 云市场镜像的specCode，即将停用。image_spec_code与image_id同时存在时取image_id的值，两者不可同时为空。
	ImageSpecCode *string `json:"image_spec_code,omitempty"`

	// 云市场镜像ID，建议使用image_id。
	ImageId *string `json:"image_id,omitempty"`
}

func (o EstimateChangeImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateChangeImageRequestBody struct{}"
	}

	return strings.Join([]string{"EstimateChangeImageRequestBody", string(data)}, " ")
}
