package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePrePaidBandwidthExtendParamOption 扩展参数，用于包周期资源变更
type UpdatePrePaidBandwidthExtendParamOption struct {

	// - 功能说明：变更资源的周期类型（包年、包月等）,可选字段。 - 取值范围：    - month-月    - year-年 - 约束：只有在资源续费降配的时候必须传，其他场景不需要传，如果传入默认忽略。
	PeriodType *UpdatePrePaidBandwidthExtendParamOptionPeriodType `json:"period_type,omitempty"`

	// - 功能说明：订购周期数，和period_type同步传入，可选字段 - 取值范围：(后续会随运营策略变化)    - period_type为month时，为[1,9]    - period_type为year时，为[1,3] - 约束：只有在资源续费降配的时候必须传，其他场景不需要传，如果传入默认忽略。该字段需要和period_type同步传入或同步不传
	PeriodNum *int32 `json:"period_num,omitempty"`

	// - 功能说明：下单订购后，是否自动从客户的账户中支付，而不需要客户手动去进行支付；系统默认是“非自动支付”。 - 取值范围：    - true：是（自动支付，从账户余额自动扣费）    - false：否（默认值，需要客户手动去支付） - 约束：资源升配或续费降配时，该参数为必传字段。自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择不自动支付，然后在用户费用中心，选择代金券支付。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o UpdatePrePaidBandwidthExtendParamOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrePaidBandwidthExtendParamOption struct{}"
	}

	return strings.Join([]string{"UpdatePrePaidBandwidthExtendParamOption", string(data)}, " ")
}

type UpdatePrePaidBandwidthExtendParamOptionPeriodType struct {
	value string
}

type UpdatePrePaidBandwidthExtendParamOptionPeriodTypeEnum struct {
	MONTH UpdatePrePaidBandwidthExtendParamOptionPeriodType
	YEAR  UpdatePrePaidBandwidthExtendParamOptionPeriodType
}

func GetUpdatePrePaidBandwidthExtendParamOptionPeriodTypeEnum() UpdatePrePaidBandwidthExtendParamOptionPeriodTypeEnum {
	return UpdatePrePaidBandwidthExtendParamOptionPeriodTypeEnum{
		MONTH: UpdatePrePaidBandwidthExtendParamOptionPeriodType{
			value: "month",
		},
		YEAR: UpdatePrePaidBandwidthExtendParamOptionPeriodType{
			value: "year",
		},
	}
}

func (c UpdatePrePaidBandwidthExtendParamOptionPeriodType) Value() string {
	return c.value
}

func (c UpdatePrePaidBandwidthExtendParamOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePrePaidBandwidthExtendParamOptionPeriodType) UnmarshalJSON(b []byte) error {
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
