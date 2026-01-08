package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChangeImageOrderRequestBody 包周期重建系统盘询价请求体。
type CreateChangeImageOrderRequestBody struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 桌面id
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *string `json:"handle_type,omitempty"`

	// 云市场镜像的specCode，即将停用。image_spec_code与image_id同时存在时取image_id的值，两者不可同时为空。
	ImageSpecCode *string `json:"image_spec_code,omitempty"`

	// 云市场镜像ID，建议使用image_id。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像类型。仅重建系统盘/更换镜像使用  - private：私有镜像。 - gold：公共镜像。
	ImageType *string `json:"image_type,omitempty"`

	// 立即重建时给用户预留的保存数据的时间（单位：分钟）。
	DelayTime *int32 `json:"delay_time,omitempty"`

	// 下发重建系统盘任务时，给用户发送的提示信息。
	Message *string `json:"message,omitempty"`
}

func (o CreateChangeImageOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChangeImageOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateChangeImageOrderRequestBody", string(data)}, " ")
}
