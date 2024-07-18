package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListP2cVgwAvailabilityZonesRequest Request Object
type ListP2cVgwAvailabilityZonesRequest struct {

	// flavor规格
	Flavor *ListP2cVgwAvailabilityZonesRequestFlavor `json:"flavor,omitempty"`
}

func (o ListP2cVgwAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVgwAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"ListP2cVgwAvailabilityZonesRequest", string(data)}, " ")
}

type ListP2cVgwAvailabilityZonesRequestFlavor struct {
	value string
}

type ListP2cVgwAvailabilityZonesRequestFlavorEnum struct {
	PROFESSIONAL1 ListP2cVgwAvailabilityZonesRequestFlavor
}

func GetListP2cVgwAvailabilityZonesRequestFlavorEnum() ListP2cVgwAvailabilityZonesRequestFlavorEnum {
	return ListP2cVgwAvailabilityZonesRequestFlavorEnum{
		PROFESSIONAL1: ListP2cVgwAvailabilityZonesRequestFlavor{
			value: "Professional1",
		},
	}
}

func (c ListP2cVgwAvailabilityZonesRequestFlavor) Value() string {
	return c.value
}

func (c ListP2cVgwAvailabilityZonesRequestFlavor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListP2cVgwAvailabilityZonesRequestFlavor) UnmarshalJSON(b []byte) error {
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
