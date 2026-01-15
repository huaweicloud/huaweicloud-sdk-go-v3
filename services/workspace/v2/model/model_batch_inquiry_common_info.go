package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchInquiryCommonInfo 批量询价/下单公共参数。
type BatchInquiryCommonInfo struct {

	// 桌面池ID。当desktop_pool_id与desktop_ids同时存在时，取desktop_ids的值，两者不可同时为空。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 包周期桌面ID列表。 不可同时存在普通桌面和池桌面ID。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 处理类型 - ONLY_FOR_EXPAND：仅对新扩容桌面生效 - FOR_EXPAND_AND_IDLE：对新扩容桌面与空闲桌面生效 - FOR_EXPAND_AND_ALL：对新扩容桌面与已有全部桌面生效
	HandleType *BatchInquiryCommonInfoHandleType `json:"handle_type,omitempty"`
}

func (o BatchInquiryCommonInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInquiryCommonInfo struct{}"
	}

	return strings.Join([]string{"BatchInquiryCommonInfo", string(data)}, " ")
}

type BatchInquiryCommonInfoHandleType struct {
	value string
}

type BatchInquiryCommonInfoHandleTypeEnum struct {
	ONLY_FOR_EXPAND     BatchInquiryCommonInfoHandleType
	FOR_EXPAND_AND_IDLE BatchInquiryCommonInfoHandleType
	FOR_EXPAND_AND_ALL  BatchInquiryCommonInfoHandleType
}

func GetBatchInquiryCommonInfoHandleTypeEnum() BatchInquiryCommonInfoHandleTypeEnum {
	return BatchInquiryCommonInfoHandleTypeEnum{
		ONLY_FOR_EXPAND: BatchInquiryCommonInfoHandleType{
			value: "ONLY_FOR_EXPAND",
		},
		FOR_EXPAND_AND_IDLE: BatchInquiryCommonInfoHandleType{
			value: "FOR_EXPAND_AND_IDLE",
		},
		FOR_EXPAND_AND_ALL: BatchInquiryCommonInfoHandleType{
			value: "FOR_EXPAND_AND_ALL",
		},
	}
}

func (c BatchInquiryCommonInfoHandleType) Value() string {
	return c.value
}

func (c BatchInquiryCommonInfoHandleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchInquiryCommonInfoHandleType) UnmarshalJSON(b []byte) error {
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
