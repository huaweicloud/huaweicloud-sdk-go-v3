package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DynamicConfigInfo 部署应用时传递的参数
type DynamicConfigInfo struct {

	// 部署应用时传递的参数名称
	Key *string `json:"key,omitempty"`

	// 部署应用时传递的参数值
	Value *string `json:"value,omitempty"`

	// 类型，如果填写动态参数，则类型必选
	Type *DynamicConfigInfoType `json:"type,omitempty"`
}

func (o DynamicConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DynamicConfigInfo struct{}"
	}

	return strings.Join([]string{"DynamicConfigInfo", string(data)}, " ")
}

type DynamicConfigInfoType struct {
	value string
}

type DynamicConfigInfoTypeEnum struct {
	TEXT       DynamicConfigInfoType
	HOST_GROUP DynamicConfigInfoType
	ENCRYPT    DynamicConfigInfoType
	ENUM       DynamicConfigInfoType
}

func GetDynamicConfigInfoTypeEnum() DynamicConfigInfoTypeEnum {
	return DynamicConfigInfoTypeEnum{
		TEXT: DynamicConfigInfoType{
			value: "text",
		},
		HOST_GROUP: DynamicConfigInfoType{
			value: "host_group",
		},
		ENCRYPT: DynamicConfigInfoType{
			value: "encrypt",
		},
		ENUM: DynamicConfigInfoType{
			value: "enum",
		},
	}
}

func (c DynamicConfigInfoType) Value() string {
	return c.value
}

func (c DynamicConfigInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DynamicConfigInfoType) UnmarshalJSON(b []byte) error {
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
