package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDeleteSubResourcesOrderRequestBody 包周期删除附属资源询价请求体。
type CreateDeleteSubResourcesOrderRequestBody struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 包周期桌面ID列表。 不可同时存在普通桌面和池桌面ID。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *CreateDeleteSubResourcesOrderRequestBodyHandleType `json:"handle_type,omitempty"`
}

func (o CreateDeleteSubResourcesOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeleteSubResourcesOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDeleteSubResourcesOrderRequestBody", string(data)}, " ")
}

type CreateDeleteSubResourcesOrderRequestBodyHandleType struct {
	value string
}

type CreateDeleteSubResourcesOrderRequestBodyHandleTypeEnum struct {
	ONLY_FOR_EXPAND     CreateDeleteSubResourcesOrderRequestBodyHandleType
	FOR_EXPAND_AND_IDLE CreateDeleteSubResourcesOrderRequestBodyHandleType
	FOR_EXPAND_AND_ALL  CreateDeleteSubResourcesOrderRequestBodyHandleType
}

func GetCreateDeleteSubResourcesOrderRequestBodyHandleTypeEnum() CreateDeleteSubResourcesOrderRequestBodyHandleTypeEnum {
	return CreateDeleteSubResourcesOrderRequestBodyHandleTypeEnum{
		ONLY_FOR_EXPAND: CreateDeleteSubResourcesOrderRequestBodyHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: CreateDeleteSubResourcesOrderRequestBodyHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: CreateDeleteSubResourcesOrderRequestBodyHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c CreateDeleteSubResourcesOrderRequestBodyHandleType) Value() string {
	return c.value
}

func (c CreateDeleteSubResourcesOrderRequestBodyHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDeleteSubResourcesOrderRequestBodyHandleType) UnmarshalJSON(b []byte) error {
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
