package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AutoRenewalResourcesReq struct {

	// |参数名称：自动续费次数| |参数的约束及描述：该参数非必填，范围限制：0-99，0代表不限制次数。 首次开通自动续费，此参数不携带或携带值为null时，默认为不限制次数 已开通自动续费，重置自动续费次数时此参数必填，否则不做处理，不进行修改|
	AutoRenewTimes *int32 `json:"auto_renew_times,omitempty"`

	// |参数名称：自动续费的周期类型| |参数的约束及描述：该参数非必填，自动续费的周期类型，支持枚举| |MONTH：包月，YEAR：包年。此参数不携带或携带值为null时，按照如下规则处理。购买时未设置自动续费功能，默认与设置资源最后一个订单的订购周期类型一致。|
	PeriodType *AutoRenewalResourcesReqPeriodType `json:"period_type,omitempty"`
}

func (o AutoRenewalResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoRenewalResourcesReq struct{}"
	}

	return strings.Join([]string{"AutoRenewalResourcesReq", string(data)}, " ")
}

type AutoRenewalResourcesReqPeriodType struct {
	value string
}

type AutoRenewalResourcesReqPeriodTypeEnum struct {
	MONTH AutoRenewalResourcesReqPeriodType
	YEAR  AutoRenewalResourcesReqPeriodType
}

func GetAutoRenewalResourcesReqPeriodTypeEnum() AutoRenewalResourcesReqPeriodTypeEnum {
	return AutoRenewalResourcesReqPeriodTypeEnum{
		MONTH: AutoRenewalResourcesReqPeriodType{
			value: "MONTH",
		},
		YEAR: AutoRenewalResourcesReqPeriodType{
			value: "YEAR",
		},
	}
}

func (c AutoRenewalResourcesReqPeriodType) Value() string {
	return c.value
}

func (c AutoRenewalResourcesReqPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoRenewalResourcesReqPeriodType) UnmarshalJSON(b []byte) error {
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
