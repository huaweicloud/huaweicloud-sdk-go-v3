package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HourPackageResource 桌面小时包资源。
type HourPackageResource struct {

	// 订购周期类型：2：月；3：年;必填
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购周期数
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动续订
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 时长用尽策略：   - SHUTDOWN_OR_HIBERNATE：自动关机/休眠。 - PAY_PER_USE：自动按需计费。
	UsedUpPolicy *HourPackageResourceUsedUpPolicy `json:"used_up_policy,omitempty"`

	// 支付后跳转url
	CloudServiceConsoleUrl *string `json:"cloud_service_console_url,omitempty"`

	CreateDesktops *CreateDesktopReq `json:"create_desktops,omitempty"`
}

func (o HourPackageResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HourPackageResource struct{}"
	}

	return strings.Join([]string{"HourPackageResource", string(data)}, " ")
}

type HourPackageResourceUsedUpPolicy struct {
	value string
}

type HourPackageResourceUsedUpPolicyEnum struct {
	SHUTDOWN_OR_HIBERNATE HourPackageResourceUsedUpPolicy
	PAY_PER_USE           HourPackageResourceUsedUpPolicy
}

func GetHourPackageResourceUsedUpPolicyEnum() HourPackageResourceUsedUpPolicyEnum {
	return HourPackageResourceUsedUpPolicyEnum{
		SHUTDOWN_OR_HIBERNATE: HourPackageResourceUsedUpPolicy{
			value: "SHUTDOWN_OR_HIBERNATE",
		},
		PAY_PER_USE: HourPackageResourceUsedUpPolicy{
			value: "PAY_PER_USE",
		},
	}
}

func (c HourPackageResourceUsedUpPolicy) Value() string {
	return c.value
}

func (c HourPackageResourceUsedUpPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HourPackageResourceUsedUpPolicy) UnmarshalJSON(b []byte) error {
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
