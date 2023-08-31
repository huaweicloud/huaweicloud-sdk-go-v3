package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateIegRequestBody struct {

	// 智能企业网关名字
	Name *string `json:"name,omitempty"`

	// 高可用性
	HighAvailability *UpdateIegRequestBodyHighAvailability `json:"high_availability,omitempty"`
}

func (o UpdateIegRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIegRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIegRequestBody", string(data)}, " ")
}

type UpdateIegRequestBodyHighAvailability struct {
	value string
}

type UpdateIegRequestBodyHighAvailabilityEnum struct {
	SINGLE UpdateIegRequestBodyHighAvailability
	DOUBLE UpdateIegRequestBodyHighAvailability
}

func GetUpdateIegRequestBodyHighAvailabilityEnum() UpdateIegRequestBodyHighAvailabilityEnum {
	return UpdateIegRequestBodyHighAvailabilityEnum{
		SINGLE: UpdateIegRequestBodyHighAvailability{
			value: "single",
		},
		DOUBLE: UpdateIegRequestBodyHighAvailability{
			value: "double",
		},
	}
}

func (c UpdateIegRequestBodyHighAvailability) Value() string {
	return c.value
}

func (c UpdateIegRequestBodyHighAvailability) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIegRequestBodyHighAvailability) UnmarshalJSON(b []byte) error {
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
