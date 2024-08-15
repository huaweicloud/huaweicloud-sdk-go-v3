package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateVgwSpecificationRequestBodyContent struct {

	// 修改后的VPN网关的规格类型
	Flavor UpdateVgwSpecificationRequestBodyContentFlavor `json:"flavor"`
}

func (o UpdateVgwSpecificationRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwSpecificationRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVgwSpecificationRequestBodyContent", string(data)}, " ")
}

type UpdateVgwSpecificationRequestBodyContentFlavor struct {
	value string
}

type UpdateVgwSpecificationRequestBodyContentFlavorEnum struct {
	BASIC                      UpdateVgwSpecificationRequestBodyContentFlavor
	PROFESSIONAL1              UpdateVgwSpecificationRequestBodyContentFlavor
	PROFESSIONAL2              UpdateVgwSpecificationRequestBodyContentFlavor
	PROFESSIONAL1_NON_FIXED_IP UpdateVgwSpecificationRequestBodyContentFlavor
	PROFESSIONAL2_NON_FIXED_IP UpdateVgwSpecificationRequestBodyContentFlavor
}

func GetUpdateVgwSpecificationRequestBodyContentFlavorEnum() UpdateVgwSpecificationRequestBodyContentFlavorEnum {
	return UpdateVgwSpecificationRequestBodyContentFlavorEnum{
		BASIC: UpdateVgwSpecificationRequestBodyContentFlavor{
			value: "Basic",
		},
		PROFESSIONAL1: UpdateVgwSpecificationRequestBodyContentFlavor{
			value: "Professional1",
		},
		PROFESSIONAL2: UpdateVgwSpecificationRequestBodyContentFlavor{
			value: "Professional2",
		},
		PROFESSIONAL1_NON_FIXED_IP: UpdateVgwSpecificationRequestBodyContentFlavor{
			value: "Professional1-NonFixedIP",
		},
		PROFESSIONAL2_NON_FIXED_IP: UpdateVgwSpecificationRequestBodyContentFlavor{
			value: "Professional2-NonFixedIP",
		},
	}
}

func (c UpdateVgwSpecificationRequestBodyContentFlavor) Value() string {
	return c.value
}

func (c UpdateVgwSpecificationRequestBodyContentFlavor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVgwSpecificationRequestBodyContentFlavor) UnmarshalJSON(b []byte) error {
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
