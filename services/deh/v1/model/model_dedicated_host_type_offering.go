package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DedicatedHostTypeOffering struct {
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Status *DedicatedHostTypeOfferingStatus `json:"status,omitempty"`
}

func (o DedicatedHostTypeOffering) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedHostTypeOffering struct{}"
	}

	return strings.Join([]string{"DedicatedHostTypeOffering", string(data)}, " ")
}

type DedicatedHostTypeOfferingStatus struct {
	value string
}

type DedicatedHostTypeOfferingStatusEnum struct {
	AVAILABLE DedicatedHostTypeOfferingStatus
	SELLOUT   DedicatedHostTypeOfferingStatus
	ABANDON   DedicatedHostTypeOfferingStatus
}

func GetDedicatedHostTypeOfferingStatusEnum() DedicatedHostTypeOfferingStatusEnum {
	return DedicatedHostTypeOfferingStatusEnum{
		AVAILABLE: DedicatedHostTypeOfferingStatus{
			value: "available",
		},
		SELLOUT: DedicatedHostTypeOfferingStatus{
			value: "sellout",
		},
		ABANDON: DedicatedHostTypeOfferingStatus{
			value: "abandon",
		},
	}
}

func (c DedicatedHostTypeOfferingStatus) Value() string {
	return c.value
}

func (c DedicatedHostTypeOfferingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DedicatedHostTypeOfferingStatus) UnmarshalJSON(b []byte) error {
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
