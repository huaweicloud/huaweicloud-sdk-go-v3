package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OrderInfoReq 云脑下单接口，包括包周期、按需场景，关联的下单场景
type OrderInfoReq struct {

	// 场景描述，执行包年包月(PREPAID)、按需(POSTPAID)开通，或者配置资源（CONFIG）用量，缺省值：PREPAID，大小写不敏感
	Scene OrderInfoReqScene `json:"scene"`

	// 促销折扣信息 String，JSON格式
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 操作类型，比如创建场景为：CREATE、订单用量预警配置为：ALERT_CONFIG，大小写不敏感
	OperateType OrderInfoReqOperateType `json:"operate_type"`

	// 计费标签
	TagList *[]SubscriptionTag `json:"tag_list,omitempty"`

	// 当scene=PREPAID 或者 POSTPAID时，当前字段必填 商品列表
	ProductList *[]Product `json:"product_list,omitempty"`

	Config *OrderConfig `json:"config,omitempty"`

	// 当scene=PREPAID时需要填写，订购周期类型： 2：月； 3：年；
	PeriodType *OrderInfoReqPeriodType `json:"period_type,omitempty"`

	// 订购周期数，当scene=PREPAID时需要填写该值 取值大于0；小于等于0会报错。 当period_type=2时，其可选范围为[1, 9]，当period_type=3，其可选范围为[1, 3]
	PeriodNum *OrderInfoReqPeriodNum `json:"period_num,omitempty"`

	// 当scene=PREPAID时，当前字段必填，是否自动续订，为空时表示不自动续订； 1：自动续订 0：不自动续订（默认）
	IsAutoRenew *OrderInfoReqIsAutoRenew `json:"is_auto_renew,omitempty"`
}

func (o OrderInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderInfoReq struct{}"
	}

	return strings.Join([]string{"OrderInfoReq", string(data)}, " ")
}

type OrderInfoReqScene struct {
	value string
}

type OrderInfoReqSceneEnum struct {
	PREPAID  OrderInfoReqScene
	POSTPAID OrderInfoReqScene
	CONFIG   OrderInfoReqScene
}

func GetOrderInfoReqSceneEnum() OrderInfoReqSceneEnum {
	return OrderInfoReqSceneEnum{
		PREPAID: OrderInfoReqScene{
			value: "PREPAID",
		},
		POSTPAID: OrderInfoReqScene{
			value: "POSTPAID",
		},
		CONFIG: OrderInfoReqScene{
			value: "CONFIG",
		},
	}
}

func (c OrderInfoReqScene) Value() string {
	return c.value
}

func (c OrderInfoReqScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrderInfoReqScene) UnmarshalJSON(b []byte) error {
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

type OrderInfoReqOperateType struct {
	value string
}

type OrderInfoReqOperateTypeEnum struct {
	CREATE       OrderInfoReqOperateType
	ALERT_CONFIG OrderInfoReqOperateType
}

func GetOrderInfoReqOperateTypeEnum() OrderInfoReqOperateTypeEnum {
	return OrderInfoReqOperateTypeEnum{
		CREATE: OrderInfoReqOperateType{
			value: "CREATE",
		},
		ALERT_CONFIG: OrderInfoReqOperateType{
			value: "ALERT_CONFIG",
		},
	}
}

func (c OrderInfoReqOperateType) Value() string {
	return c.value
}

func (c OrderInfoReqOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrderInfoReqOperateType) UnmarshalJSON(b []byte) error {
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

type OrderInfoReqPeriodType struct {
	value int32
}

type OrderInfoReqPeriodTypeEnum struct {
	E_2 OrderInfoReqPeriodType
	E_3 OrderInfoReqPeriodType
}

func GetOrderInfoReqPeriodTypeEnum() OrderInfoReqPeriodTypeEnum {
	return OrderInfoReqPeriodTypeEnum{
		E_2: OrderInfoReqPeriodType{
			value: 2,
		}, E_3: OrderInfoReqPeriodType{
			value: 3,
		},
	}
}

func (c OrderInfoReqPeriodType) Value() int32 {
	return c.value
}

func (c OrderInfoReqPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrderInfoReqPeriodType) UnmarshalJSON(b []byte) error {
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

type OrderInfoReqPeriodNum struct {
	value int32
}

type OrderInfoReqPeriodNumEnum struct {
	E_1 OrderInfoReqPeriodNum
	E_2 OrderInfoReqPeriodNum
	E_3 OrderInfoReqPeriodNum
	E_4 OrderInfoReqPeriodNum
	E_5 OrderInfoReqPeriodNum
	E_6 OrderInfoReqPeriodNum
	E_7 OrderInfoReqPeriodNum
	E_8 OrderInfoReqPeriodNum
	E_9 OrderInfoReqPeriodNum
}

func GetOrderInfoReqPeriodNumEnum() OrderInfoReqPeriodNumEnum {
	return OrderInfoReqPeriodNumEnum{
		E_1: OrderInfoReqPeriodNum{
			value: 1,
		}, E_2: OrderInfoReqPeriodNum{
			value: 2,
		}, E_3: OrderInfoReqPeriodNum{
			value: 3,
		}, E_4: OrderInfoReqPeriodNum{
			value: 4,
		}, E_5: OrderInfoReqPeriodNum{
			value: 5,
		}, E_6: OrderInfoReqPeriodNum{
			value: 6,
		}, E_7: OrderInfoReqPeriodNum{
			value: 7,
		}, E_8: OrderInfoReqPeriodNum{
			value: 8,
		}, E_9: OrderInfoReqPeriodNum{
			value: 9,
		},
	}
}

func (c OrderInfoReqPeriodNum) Value() int32 {
	return c.value
}

func (c OrderInfoReqPeriodNum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrderInfoReqPeriodNum) UnmarshalJSON(b []byte) error {
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

type OrderInfoReqIsAutoRenew struct {
	value int32
}

type OrderInfoReqIsAutoRenewEnum struct {
	E_0 OrderInfoReqIsAutoRenew
	E_1 OrderInfoReqIsAutoRenew
}

func GetOrderInfoReqIsAutoRenewEnum() OrderInfoReqIsAutoRenewEnum {
	return OrderInfoReqIsAutoRenewEnum{
		E_0: OrderInfoReqIsAutoRenew{
			value: 0,
		}, E_1: OrderInfoReqIsAutoRenew{
			value: 1,
		},
	}
}

func (c OrderInfoReqIsAutoRenew) Value() int32 {
	return c.value
}

func (c OrderInfoReqIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrderInfoReqIsAutoRenew) UnmarshalJSON(b []byte) error {
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
