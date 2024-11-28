package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateCbcOrderRequestBody struct {

	// 计费模式： 0：一次性计费（默认值，对应包年包月） 10：RI
	ChargingMode CreateCbcOrderRequestBodyChargingMode `json:"charging_mode"`

	// 0：不自动续订 1：自动续订
	IsAutoRenew CreateCbcOrderRequestBodyIsAutoRenew `json:"is_auto_renew"`

	// 该请求参数为预留参数，当前不支持自动支付，使用接口时该参数请使用0 0：不自动支付 1：自动支付
	IsAutoPay CreateCbcOrderRequestBodyIsAutoPay `json:"is_auto_pay"`

	// period_num
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 订购周期类型： 2：月； 3：年； 4：包小时（仅限带宽加油包购买场景使用） 5：绝对时间；（追加附属资源场景使用，比如主机上追加云硬盘） 6：一次性（chargingMode=1 一次性计费场景使用）
	PeriodType int32 `json:"period_type"`

	// 用户购买的云服务的主云服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// project_id
	ProjectId string `json:"project_id"`

	// promotion_info
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// Region标识，填region编码如\"cn-north-1\"，对于global服务，此处固定填写虚拟的Global regionCode(global-cbc-1)
	RegionId string `json:"region_id"`

	// product_infos
	ProductInfos []ProductInfo `json:"product_infos"`
}

func (o CreateCbcOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCbcOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCbcOrderRequestBody", string(data)}, " ")
}

type CreateCbcOrderRequestBodyChargingMode struct {
	value int32
}

type CreateCbcOrderRequestBodyChargingModeEnum struct {
	E_1  CreateCbcOrderRequestBodyChargingMode
	E_10 CreateCbcOrderRequestBodyChargingMode
}

func GetCreateCbcOrderRequestBodyChargingModeEnum() CreateCbcOrderRequestBodyChargingModeEnum {
	return CreateCbcOrderRequestBodyChargingModeEnum{
		E_1: CreateCbcOrderRequestBodyChargingMode{
			value: 1,
		}, E_10: CreateCbcOrderRequestBodyChargingMode{
			value: 10,
		},
	}
}

func (c CreateCbcOrderRequestBodyChargingMode) Value() int32 {
	return c.value
}

func (c CreateCbcOrderRequestBodyChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCbcOrderRequestBodyChargingMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateCbcOrderRequestBodyIsAutoRenew struct {
	value int32
}

type CreateCbcOrderRequestBodyIsAutoRenewEnum struct {
	E_0 CreateCbcOrderRequestBodyIsAutoRenew
	E_1 CreateCbcOrderRequestBodyIsAutoRenew
}

func GetCreateCbcOrderRequestBodyIsAutoRenewEnum() CreateCbcOrderRequestBodyIsAutoRenewEnum {
	return CreateCbcOrderRequestBodyIsAutoRenewEnum{
		E_0: CreateCbcOrderRequestBodyIsAutoRenew{
			value: 0,
		}, E_1: CreateCbcOrderRequestBodyIsAutoRenew{
			value: 1,
		},
	}
}

func (c CreateCbcOrderRequestBodyIsAutoRenew) Value() int32 {
	return c.value
}

func (c CreateCbcOrderRequestBodyIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCbcOrderRequestBodyIsAutoRenew) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateCbcOrderRequestBodyIsAutoPay struct {
	value int32
}

type CreateCbcOrderRequestBodyIsAutoPayEnum struct {
	E_0 CreateCbcOrderRequestBodyIsAutoPay
	E_1 CreateCbcOrderRequestBodyIsAutoPay
}

func GetCreateCbcOrderRequestBodyIsAutoPayEnum() CreateCbcOrderRequestBodyIsAutoPayEnum {
	return CreateCbcOrderRequestBodyIsAutoPayEnum{
		E_0: CreateCbcOrderRequestBodyIsAutoPay{
			value: 0,
		}, E_1: CreateCbcOrderRequestBodyIsAutoPay{
			value: 1,
		},
	}
}

func (c CreateCbcOrderRequestBodyIsAutoPay) Value() int32 {
	return c.value
}

func (c CreateCbcOrderRequestBodyIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCbcOrderRequestBodyIsAutoPay) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
