package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *EstimateChangeImageRequestBodyHandleType `json:"handle_type,omitempty"`

	// 云市场镜像的specCode，即将停用。image_spec_code与image_id同时存在时取image_id的值，两者不可同时为空。
	ImageSpecCode *string `json:"image_spec_code,omitempty"`

	// 云市场镜像ID，建议使用image_id。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像类型。仅重建系统盘/更换镜像使用  - private：私有镜像。 - gold：公共镜像。
	ImageType *string `json:"image_type,omitempty"`
}

func (o EstimateChangeImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateChangeImageRequestBody struct{}"
	}

	return strings.Join([]string{"EstimateChangeImageRequestBody", string(data)}, " ")
}

type EstimateChangeImageRequestBodyHandleType struct {
	value string
}

type EstimateChangeImageRequestBodyHandleTypeEnum struct {
	ONLY_FOR_EXPAND     EstimateChangeImageRequestBodyHandleType
	FOR_EXPAND_AND_IDLE EstimateChangeImageRequestBodyHandleType
	FOR_EXPAND_AND_ALL  EstimateChangeImageRequestBodyHandleType
}

func GetEstimateChangeImageRequestBodyHandleTypeEnum() EstimateChangeImageRequestBodyHandleTypeEnum {
	return EstimateChangeImageRequestBodyHandleTypeEnum{
		ONLY_FOR_EXPAND: EstimateChangeImageRequestBodyHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: EstimateChangeImageRequestBodyHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: EstimateChangeImageRequestBodyHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c EstimateChangeImageRequestBodyHandleType) Value() string {
	return c.value
}

func (c EstimateChangeImageRequestBodyHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EstimateChangeImageRequestBodyHandleType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
