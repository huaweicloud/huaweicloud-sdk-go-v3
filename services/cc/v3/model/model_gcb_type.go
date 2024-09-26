package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GcbType 全域互联带宽类型。
type GcbType struct {

	// 功能说明：描述带宽类型，对应地理区间的城域、区域、大区、跨区四级： - TrsArea: 跨区带宽 - Area: 大区带宽 - SubArea: 区域带宽 - Region: 城域带宽
	Type GcbTypeType `json:"type"`
}

func (o GcbType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbType struct{}"
	}

	return strings.Join([]string{"GcbType", string(data)}, " ")
}

type GcbTypeType struct {
	value string
}

type GcbTypeTypeEnum struct {
	REGION   GcbTypeType
	SUB_AREA GcbTypeType
	AREA     GcbTypeType
	TRS_AREA GcbTypeType
}

func GetGcbTypeTypeEnum() GcbTypeTypeEnum {
	return GcbTypeTypeEnum{
		REGION: GcbTypeType{
			value: "Region",
		},
		SUB_AREA: GcbTypeType{
			value: "SubArea",
		},
		AREA: GcbTypeType{
			value: "Area",
		},
		TRS_AREA: GcbTypeType{
			value: "TrsArea",
		},
	}
}

func (c GcbTypeType) Value() string {
	return c.value
}

func (c GcbTypeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GcbTypeType) UnmarshalJSON(b []byte) error {
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
