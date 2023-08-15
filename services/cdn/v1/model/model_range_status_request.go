package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RangeStatusRequest struct {

	// range状态（\"off\"/\"on\"）
	RangeStatus RangeStatusRequestRangeStatus `json:"range_status"`
}

func (o RangeStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RangeStatusRequest struct{}"
	}

	return strings.Join([]string{"RangeStatusRequest", string(data)}, " ")
}

type RangeStatusRequestRangeStatus struct {
	value string
}

type RangeStatusRequestRangeStatusEnum struct {
	OFF RangeStatusRequestRangeStatus
	ON  RangeStatusRequestRangeStatus
}

func GetRangeStatusRequestRangeStatusEnum() RangeStatusRequestRangeStatusEnum {
	return RangeStatusRequestRangeStatusEnum{
		OFF: RangeStatusRequestRangeStatus{
			value: "off",
		},
		ON: RangeStatusRequestRangeStatus{
			value: "on",
		},
	}
}

func (c RangeStatusRequestRangeStatus) Value() string {
	return c.value
}

func (c RangeStatusRequestRangeStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RangeStatusRequestRangeStatus) UnmarshalJSON(b []byte) error {
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
