package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BindingGeipBody struct {

	// 全局弹性公网IP的ID
	GlobalEipId string `json:"global_eip_id"`

	// geip子网类型
	Type *BindingGeipBodyType `json:"type,omitempty"`
}

func (o BindingGeipBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingGeipBody struct{}"
	}

	return strings.Join([]string{"BindingGeipBody", string(data)}, " ")
}

type BindingGeipBodyType struct {
	value string
}

type BindingGeipBodyTypeEnum struct {
	IP_ADDRESS BindingGeipBodyType
	IP_SEGMENT BindingGeipBodyType
}

func GetBindingGeipBodyTypeEnum() BindingGeipBodyTypeEnum {
	return BindingGeipBodyTypeEnum{
		IP_ADDRESS: BindingGeipBodyType{
			value: "IP_ADDRESS",
		},
		IP_SEGMENT: BindingGeipBodyType{
			value: "IP_SEGMENT",
		},
	}
}

func (c BindingGeipBodyType) Value() string {
	return c.value
}

func (c BindingGeipBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BindingGeipBodyType) UnmarshalJSON(b []byte) error {
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
