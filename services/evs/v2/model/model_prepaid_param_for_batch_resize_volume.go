package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrepaidParamForBatchResizeVolume The billing policy parameter for the yearly/monthly disk capacity expansion.
type PrepaidParamForBatchResizeVolume struct {

	// Whether to pay immediately. This parameter is valid only when the disk is billed on a yearly/monthly basis. The default value is **false**. Values: - **true**: An order is immediately paid from the account balance. - **false**: An order is not paid immediately after being created.
	IsAutoPay *PrepaidParamForBatchResizeVolumeIsAutoPay `json:"is_auto_pay,omitempty"`
}

func (o PrepaidParamForBatchResizeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidParamForBatchResizeVolume struct{}"
	}

	return strings.Join([]string{"PrepaidParamForBatchResizeVolume", string(data)}, " ")
}

type PrepaidParamForBatchResizeVolumeIsAutoPay struct {
	value string
}

type PrepaidParamForBatchResizeVolumeIsAutoPayEnum struct {
	FALSE PrepaidParamForBatchResizeVolumeIsAutoPay
	TRUE  PrepaidParamForBatchResizeVolumeIsAutoPay
}

func GetPrepaidParamForBatchResizeVolumeIsAutoPayEnum() PrepaidParamForBatchResizeVolumeIsAutoPayEnum {
	return PrepaidParamForBatchResizeVolumeIsAutoPayEnum{
		FALSE: PrepaidParamForBatchResizeVolumeIsAutoPay{
			value: "false",
		},
		TRUE: PrepaidParamForBatchResizeVolumeIsAutoPay{
			value: "true",
		},
	}
}

func (c PrepaidParamForBatchResizeVolumeIsAutoPay) Value() string {
	return c.value
}

func (c PrepaidParamForBatchResizeVolumeIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidParamForBatchResizeVolumeIsAutoPay) UnmarshalJSON(b []byte) error {
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
