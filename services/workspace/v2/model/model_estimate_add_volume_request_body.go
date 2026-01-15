package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EstimateAddVolumeRequestBody 包周期桌面池添加磁盘、切换镜像询价请求体。
type EstimateAddVolumeRequestBody struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 包周期桌面ID列表。 不可同时存在普通桌面和池桌面ID。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *EstimateAddVolumeRequestBodyHandleType `json:"handle_type,omitempty"`
}

func (o EstimateAddVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateAddVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"EstimateAddVolumeRequestBody", string(data)}, " ")
}

type EstimateAddVolumeRequestBodyHandleType struct {
	value string
}

type EstimateAddVolumeRequestBodyHandleTypeEnum struct {
	ONLY_FOR_EXPAND     EstimateAddVolumeRequestBodyHandleType
	FOR_EXPAND_AND_IDLE EstimateAddVolumeRequestBodyHandleType
	FOR_EXPAND_AND_ALL  EstimateAddVolumeRequestBodyHandleType
}

func GetEstimateAddVolumeRequestBodyHandleTypeEnum() EstimateAddVolumeRequestBodyHandleTypeEnum {
	return EstimateAddVolumeRequestBodyHandleTypeEnum{
		ONLY_FOR_EXPAND: EstimateAddVolumeRequestBodyHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: EstimateAddVolumeRequestBodyHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: EstimateAddVolumeRequestBodyHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c EstimateAddVolumeRequestBodyHandleType) Value() string {
	return c.value
}

func (c EstimateAddVolumeRequestBodyHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EstimateAddVolumeRequestBodyHandleType) UnmarshalJSON(b []byte) error {
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
