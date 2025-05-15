package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateBandwidthPackage 更新带宽包实例的请求体
type UpdateBandwidthPackage struct {

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 带宽包实例中的带宽值。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 带宽包实例在大陆站或国际站的计费方式： - 5：大陆站按95方式计费 - 6：国际站按95方式计费
	BillingMode *UpdateBandwidthPackageBillingMode `json:"billing_mode,omitempty"`
}

func (o UpdateBandwidthPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthPackage struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthPackage", string(data)}, " ")
}

type UpdateBandwidthPackageBillingMode struct {
	value int32
}

type UpdateBandwidthPackageBillingModeEnum struct {
	E_5 UpdateBandwidthPackageBillingMode
	E_6 UpdateBandwidthPackageBillingMode
}

func GetUpdateBandwidthPackageBillingModeEnum() UpdateBandwidthPackageBillingModeEnum {
	return UpdateBandwidthPackageBillingModeEnum{
		E_5: UpdateBandwidthPackageBillingMode{
			value: 5,
		}, E_6: UpdateBandwidthPackageBillingMode{
			value: 6,
		},
	}
}

func (c UpdateBandwidthPackageBillingMode) Value() int32 {
	return c.value
}

func (c UpdateBandwidthPackageBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateBandwidthPackageBillingMode) UnmarshalJSON(b []byte) error {
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
