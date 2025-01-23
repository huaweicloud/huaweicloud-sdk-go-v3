package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradePrepaidOption 参数解释：创建负载均衡器实例的预付费计费配置。若传入该结构体，则创建预付费类型的负载均衡器实例。  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,fcs,dt,hcso_dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,srg)
type UpgradePrepaidOption struct {

	// 参数解释：预付费实例的订购周期类型，当前支持月和年。  取值范围：  - month：月。  - year：年。
	PeriodType UpgradePrepaidOptionPeriodType `json:"period_type"`

	// 参数解释：预付费实例的订购周期数。  取值范围： - period_type为month时，为[1,9]。 - period_type为year时，为[1,3]。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 参数解释：购买定向套餐包。
	ResourcePackageType []string `json:"resource_package_type"`

	// 参数解释：自动支付开关。下单订购后，是否自动从客户的账户中支付。  约束限制：开启自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择关闭自动支付，然后在用户费用中心，选择代金券支付。  取值范围：  - true：开启自动支付。  - false：关闭自动支付。
	AutoPay *bool `json:"auto_pay,omitempty"`
}

func (o UpgradePrepaidOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradePrepaidOption struct{}"
	}

	return strings.Join([]string{"UpgradePrepaidOption", string(data)}, " ")
}

type UpgradePrepaidOptionPeriodType struct {
	value string
}

type UpgradePrepaidOptionPeriodTypeEnum struct {
	MONTH UpgradePrepaidOptionPeriodType
	YEAR  UpgradePrepaidOptionPeriodType
}

func GetUpgradePrepaidOptionPeriodTypeEnum() UpgradePrepaidOptionPeriodTypeEnum {
	return UpgradePrepaidOptionPeriodTypeEnum{
		MONTH: UpgradePrepaidOptionPeriodType{
			value: "month",
		},
		YEAR: UpgradePrepaidOptionPeriodType{
			value: "year",
		},
	}
}

func (c UpgradePrepaidOptionPeriodType) Value() string {
	return c.value
}

func (c UpgradePrepaidOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradePrepaidOptionPeriodType) UnmarshalJSON(b []byte) error {
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
